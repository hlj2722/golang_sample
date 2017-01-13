package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "hlj2722"
	fmt.Println("原字符串:",data)
	// 编码
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("编码:",sEnc)
	// 解码
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("解码:",string(sDec))
	fmt.Println()
}
