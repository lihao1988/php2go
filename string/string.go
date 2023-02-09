package string

import (
	"strings"
	"unicode/utf8"
)

// StrLen string length
// php strlen
func StrLen(s string) int {
	return len(s)
}

// MbStrLen string for utf8
// php mb_strlen
func MbStrLen(s string) int {
	return utf8.RuneCountInString(s)
}

// SubstrCount counts the number of substr in s
// php substr_count
func SubstrCount(s, substr string) int {
	return strings.Count(s, substr)
}

// Substr get substr
// php substr
func Substr(s string, start int, length int) string {
	return s[start : start+length]
}

// MbSubstr get substr for utf8
// php mb_substr
func MbSubstr(s string, start int, length int) string {
	strRune := []rune(s)
	return string(strRune[start : start+length])
}

// StrPos get first index of substr in s, from "start" index
// php strpos
func StrPos(s, substr string, start int) int {
	return strings.Index(s[start:], substr)
}

// StrRPos get last index of substr in s, from "start" index
// php strrpos
func StrRPos(s, substr string, start int) int {
	return strings.LastIndex(s[start:], substr)
}

// StrSplit slices s into all substrings separated by sep
// php str_split
func StrSplit(s, sep string) []string {
	return strings.Split(s, sep)
}

// UCFirst toUpper the first letter
// php ucfirst
func UCFirst(s string) string {
	firstChart := strings.ToUpper(s[0:1])
	return firstChart + s[1:]
}
