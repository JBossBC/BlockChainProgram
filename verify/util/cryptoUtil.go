package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5Crypto(data string) string {
	hash := md5.New()
	io.WriteString(hash, data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
