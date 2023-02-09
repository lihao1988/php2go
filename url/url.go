package url

import (
	netUrl "net/url"
	"strings"
)

// ParseUrl parse url_string
// php parse_url
func ParseUrl(rawURL string) (*netUrl.URL, error) {
	return netUrl.Parse(rawURL)
}

// UrlEncode url-encode string
// php urlencode
func UrlEncode(str string) string {
	return netUrl.QueryEscape(str)
}

// UrlDecode url-decode string
// php urldecode
func UrlDecode(str string) (string, error) {
	return netUrl.QueryUnescape(str)
}

// RawUrlEncode raw url-encode string
// php rawurlencode
func RawUrlEncode(str string) string {
	return strings.Replace(netUrl.QueryEscape(str), "+", "%20", -1)
}

// RawUrlDecode raw url-decode string
// php rawurldecode
func RawUrlDecode(str string) (string, error) {
	return netUrl.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}

// HttpBuildQuery http build query
// php http_build_query
func HttpBuildQuery(queryData map[string]string) string {
	var uri netUrl.URL
	q := uri.Query()
	for key, value := range queryData {
		if strings.TrimSpace(value) != "" {
			q.Add(key, value)
		}
	}

	return q.Encode()
}
