package utils

import (
	"encoding/base64"
	"math/rand"
	"regexp"
)

var reConvertChars = regexp.MustCompile(`[\/\+]`)

func RandomCode(n ...int) string {
	size := 12
	if len(n) > 0 && n[0]%3 == 0 {
		size = n[0]
	}
	buf := make([]byte, size)
	rand.Read(buf)
	return reConvertChars.ReplaceAllString(base64.StdEncoding.EncodeToString(buf), "0")
}
