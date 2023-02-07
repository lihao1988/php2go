package crypto

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	str := "你好啊"

	fmt.Println("MD5: ", MD5(str))
	fmt.Println("Sha1: ", Sha1(str))

	base64Str := Base64Encode(str)
	fmt.Println("Base64Encode: ", base64Str)
	bStr, err := Base64Decode(base64Str)
	fmt.Println("Base64Decode: ", bStr, ", Error: ", err)

	urlStr := UrlEncode(str)
	fmt.Println("UrlEncode: ", urlStr)
	uStr, err := UrlDecode(base64Str)
	fmt.Println("UrlDecode: ", uStr, ", Error: ", err)
}
