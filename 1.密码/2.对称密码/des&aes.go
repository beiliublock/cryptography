package main

import (
	"crypto/aes"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func main() {
	//注意：DES和AES底层都有使用到异或，所以明文必须为分组长度，不然会报错
	//可以通过添加填充来满足条件
	//DES不用分组模式迭代的话一次只能加密64bit的明文
	block, _ := des.NewCipher([]byte("12345678"))
	dst := make([]byte,des.BlockSize)
	src := make([]byte,des.BlockSize)
	//加密
	block.Encrypt(dst,[]byte("87654321"))
	s := hex.EncodeToString(dst)
	fmt.Println("密文："+s)
	//解密
	bytes, _ := hex.DecodeString(s)
	block.Decrypt(src,bytes)
	fmt.Printf("明文：%s\n",src)

	//同理，AES不用分组模式迭代的话一次只能加密128bit的明文
	cipher, _ := aes.NewCipher([]byte("1234567887654321"))
	dst = make([]byte,aes.BlockSize)
	src = make([]byte,aes.BlockSize)
	//加密
	cipher.Encrypt(dst,[]byte("abcdefgh98765432"))
	s = hex.EncodeToString(dst)
	fmt.Println("密文："+s)
	//解密
	bytes, _ = hex.DecodeString(s)
	cipher.Decrypt(src,bytes)
	fmt.Printf("明文：%s",src)

}
