package url

import (
	"fmt"
	"testing"
)

func TestUrl(t *testing.T) {
	str := "你好啊~"

	urlStr := UrlEncode(str)
	fmt.Println("UrlEncode: ", urlStr)
	uStr, err := UrlDecode(urlStr)
	fmt.Println("UrlDecode: ", uStr, ", Error: ", err)
}
