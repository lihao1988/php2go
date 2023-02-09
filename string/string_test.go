package string

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	str := "中国88-china-88"
	fmt.Println("StrLen: ", StrLen(str))
	fmt.Println("MbStrLen: ", MbStrLen(str))

	haystack := "I love Shanghai. Shanghai is the biggest city in china."
	needle := "Shanghai"
	fmt.Println("SubstrCount: ", SubstrCount(haystack, needle))

	fmt.Println("Substr: ", Substr(str, 1, 2))
	fmt.Println("MbSubstr: ", MbSubstr(str, 1, 2))

	fmt.Println("StrPos: ", StrPos(str, "-", 2))
	fmt.Println("StrRPos: ", StrRPos(str, "-", 9))

	fmt.Println("UCFirst: ", UCFirst("abcdefg"))
}
