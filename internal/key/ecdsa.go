package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func GenerateEcdsaKeys() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return &privateKey.PublicKey, privateKey
}
