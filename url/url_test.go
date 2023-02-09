package url

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUrl(t *testing.T) {
	urlStr := "https://www.example.com/path?googleguy=googley"
	urlData, err := ParseUrl(urlStr)
	urlDataByte, _ := json.Marshal(urlData)
	fmt.Println("ParseUrl: ", string(urlDataByte), ", Error: ", err)

	str := "你好 啊~ "
	urlEnStr := UrlEncode(str)
	fmt.Println("UrlEncode: ", urlEnStr)
	uStr, err := UrlDecode(urlEnStr)
	fmt.Println("UrlDecode: ", uStr, ", Error: ", err)

	urlREnStr := RawUrlEncode(str)
	fmt.Println("RawUrlEncode: ", urlREnStr)
	uRStr, err := RawUrlDecode(urlREnStr)
	fmt.Println("RawUrlDecode: ", uRStr, ", Error: ", err)

	queryData := map[string]string{
		"a": "a1", "b": "b1",
	}
	fmt.Println("HttpBuildQuery: ", HttpBuildQuery(queryData))
}
