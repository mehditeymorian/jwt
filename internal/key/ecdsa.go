package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func GenerateEcdsaKeys(curve string) (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	var ellipticCurve elliptic.Curve

	switch curve {
	case "P224":
		ellipticCurve = elliptic.P224()
	case "P384":
		ellipticCurve = elliptic.P384()
	case "P521":
		ellipticCurve = elliptic.P521()
	case "P256":
		fallthrough
	default:
		ellipticCurve = elliptic.P256()
	}
	privateKey, err := ecdsa.GenerateKey(ellipticCurve, rand.Reader)
	if err != nil {
		return nil, nil
	}

	return &privateKey.PublicKey, privateKey
}
