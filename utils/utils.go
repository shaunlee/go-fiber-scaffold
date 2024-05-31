package utils

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

func RandomCode(n ...int) string {
	size := 12
	if len(n) > 0 && n[0] > 2 {
		size = n[0] - n[0]%3
	}
	buf := make([]byte, size)
	rand.Read(buf)
	return strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(buf), "/", "0"), "+", "0")
}
