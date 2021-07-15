package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(passcode string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(passcode))
	return hex.EncodeToString(hash.Sum(nil))

}
