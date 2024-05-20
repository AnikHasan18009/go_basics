package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Hash(s string) string {
	hash := sha1.Sum([]byte(s))

	hashString := hex.EncodeToString(hash[:])
	return hashString

}
