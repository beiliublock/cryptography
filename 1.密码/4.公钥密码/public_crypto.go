package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	privateKey, e := ioutil.ReadFile(`src\密码学\1.密码\4.公钥密码\private.pem`)
	if e != nil {
		fmt.Println("私钥获取失败")
	}
	publicKey, i := ioutil.ReadFile(`src\密码学\1.密码\4.公钥密码\public.pem`)
	if i != nil {
		fmt.Println("公钥获取失败")
	}
	fmt.Printf("%x\n",privateKey)
	fmt.Printf("%x",publicKey)

}

//生成公私钥，但是生产环境往往是文件中读取
//通常直接用openssl生成
//私钥：openssl genrsa -out private.pem 1024
//公钥：openssl rsa -in private.pem -pubout -out public.pem
func generateKey(num int) (*rsa.PrivateKey,*rsa.PublicKey,error) {
	privateKey, e := rsa.GenerateKey(rand.Reader, num)
	if e !=nil {
		return nil,nil,errors.New("密钥生成失败")
	}
	publicKey := privateKey.PublicKey
	return privateKey,&publicKey,nil
}

// 使用对方的公钥的数据, 只有对方的私钥才能解开
func encrypt(plain string, publicKey string) (cipherByte []byte, err error) {
	msg := []byte(plain)
	// 解码公钥
	pubBlock, _ := pem.Decode([]byte(publicKey))
	// 读取公钥
	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(err)
	}
	pub := pubKeyValue.(*rsa.PublicKey)
	// 加密数据方法: 不用使用EncryptPKCS1v15方法加密,源码里面推荐使用EncryptOAEP, 因此这里使用安全的方法加密
	encryptOAEP, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, msg, nil)
	if err != nil {
		panic(err)
	}
	cipherByte = encryptOAEP
	return
}

// 使用私钥解密公钥加密的数据
func decrypt(cipherByte []byte, privateKey string) (plainText string, err error) {
	// 解析出私钥
	priBlock, _ := pem.Decode([]byte(privateKey))
	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密RSA-OAEP方式加密后的内容
	decryptOAEP, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		panic(err)
	}
	plainText = string(decryptOAEP)
	return
}