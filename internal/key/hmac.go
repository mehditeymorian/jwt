package key

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateHmacKey(base64Encoded bool) []byte {
	size := 256

	key := make([]byte, size)

	_, err := rand.Reader.Read(key)
	if err != nil {
		return nil
	}

	if base64Encoded {
		encodedKey := make([]byte, size)

		base64.StdEncoding.Encode(encodedKey, key)

		return encodedKey
	}

	return key
}
