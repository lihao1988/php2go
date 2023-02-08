package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MD5 encode string to md5 string
// php md5
func MD5(str string) string {
	data := md5.Sum([]byte(str))

	// 将 []byte 转换为 32字符十六进制数
	return fmt.Sprintf("%x", data)
}

// Sha1 encode string to sha1 string
// php sha1
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Base64Encode encode string to base64 string
// php base64_encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode decode base64 string to normal string
// php base64_decode
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}
