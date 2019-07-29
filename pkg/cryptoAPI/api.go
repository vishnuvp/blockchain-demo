package cryptoAPI

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256Hash(data string) string{
	hash := sha256.New()
	hash.Write([]byte(data))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}
