package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func createHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateShortURL(url string) string {
	finalString := base64.StdEncoding.EncodeToString([]byte(createHash(url)))
	if len(finalString) > 10 {
		return finalString[:10]
	}
	return finalString
}
