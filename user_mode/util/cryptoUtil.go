package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func MD5Crypto(data string) string {
	hash := md5.New()
	io.WriteString(hash, data)
	var cryptoData = fmt.Sprintf("%x", hash.Sum(nil))

	return hex.EncodeToString([]byte(cryptoData))
}
