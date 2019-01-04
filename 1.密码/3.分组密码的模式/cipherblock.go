package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

//这里只拿DES加密算法中的CBC(密文分组链接模式)举例，其他的分组密码和分组密码模式一样
//golang中好像没有提供ECB(电子密码本模式)，所以基本所有的分组都需要初始化向量(IV)
func main() {
	block, _ := des.NewCipher([]byte("abcdefgh"))
	//填充数据,进行加密
	text, l := paddingText("9")
	dst := make([]byte, len(text))
	cipher.NewCBCEncrypter(block, []byte("12345678")).CryptBlocks(dst, text)
	s := hex.EncodeToString(dst)
	fmt.Println("密文：", s)
	//解密后，将填充的数据减去
	src := make([]byte, len(text))
	cipher.NewCBCDecrypter(block,[]byte("12345678")).CryptBlocks(src,dst)
	i := trimText(src, l)
	fmt.Printf("%s",i)
}

//填充数据
func paddingText(str string) ([]byte, int) {
	src := []byte(str)
	l := des.BlockSize - len(src)%des.BlockSize
	padText := bytes.Repeat([]byte{0}, l)
	src = append(src, padText...)
	return src, l
}

//解密后，将填充的数据减去
func trimText(src []byte,l int) []byte {
	trimText := bytes.Repeat([]byte{0}, l)
	right := bytes.TrimRight(src, string(trimText))
	return right
}
