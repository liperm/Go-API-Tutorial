package encryptions

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncriptData(data string) string {
	encryption := sha256.New()
	encryption.Write([]byte(data))
	return hex.EncodeToString(encryption.Sum(nil))
}
