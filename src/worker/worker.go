package worker

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Match(message string, target string) (bool, string) {
	sum := md5Hash(message)

	if strings.Compare(target, sum) == 0 {
		return true, message
	}

	return false, ""
}
