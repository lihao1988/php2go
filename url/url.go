package url

import "net/url"

// UrlEncode url-encode string
// php urlencode
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// UrlDecode url-decode string
// php urldecode
func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
