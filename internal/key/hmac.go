package key

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/pterm/pterm"
)

func GenerateHmacKey(size int, base64Encoded bool) []byte {

	key := make([]byte, size)

	_, err := rand.Reader.Read(key)
	if err != nil {
		pterm.Fatal.Println("failed to generate hmac key")

		return nil
	}

	if base64Encoded {
		encodedKey := make([]byte, base64.StdEncoding.EncodedLen(size))

		base64.StdEncoding.Encode(encodedKey, key)

		return encodedKey
	}

	return key
}
