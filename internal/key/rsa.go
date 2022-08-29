package key

import (
	"crypto/rand"
	"crypto/rsa"
)

func GenerateRsaKeys() (*rsa.PublicKey, *rsa.PrivateKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil
	}

	return &privateKey.PublicKey, privateKey
}
