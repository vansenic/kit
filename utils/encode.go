package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

func CreateDigest(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64EncodeSafe - Escape before encode
func Base64EncodeSafe(v []byte) string {
	temp := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(string(v))))
	return temp
}

// Base64DecodeSafe - Unescape after decode
func Base64DecodeSafe(v string) (string, error) {
	temp, err := base64.StdEncoding.DecodeString(v)
	t, _ := url.QueryUnescape(string(temp))
	return t, err
}

func Base64Encode(v []byte) string {
	temp := base64.StdEncoding.EncodeToString(v)
	return temp
}

func Base64Decode(v string) (string, error) {
	temp, err := base64.StdEncoding.DecodeString(v)
	return string(temp), err
}
