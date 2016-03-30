package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	token = "a07fbd22a3bd11e5ada43010b3a0f"
)

func SignValidation(signature, timestamp, nonce string) bool {
	var array = []string{token, timestamp, nonce}
	sort.Strings(array)
	s1 := sha1.New()
	io.WriteString(s1, strings.Join(array, ""))
	sign := fmt.Sprintf("%x", s1.Sum(nil))
	return sign == signature
}
