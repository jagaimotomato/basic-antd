package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

/* Md5 hash */
func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
