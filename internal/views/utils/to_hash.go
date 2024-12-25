package utils

import (
	"crypto/sha256"
	"fmt"
)

func ToHash(s string) string {
	hash := sha256.Sum256([]byte(s))
	hash_s := fmt.Sprintf("%x", hash)
	return hash_s
}
