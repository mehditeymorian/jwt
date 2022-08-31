package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"github.com/pterm/pterm"
)

func GenerateEcdsaKeys(curve string) (string, string) {
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
		pterm.Warning.Println("elliptic curve is not valid. choosing P256 as elliptic curve")

		ellipticCurve = elliptic.P256()
	}
	privateKey, err := ecdsa.GenerateKey(ellipticCurve, rand.Reader)
	if err != nil {
		pterm.Fatal.Println("failed to generate ecdsa keys")

		return "", ""
	}

	publicKey := &privateKey.PublicKey

	publicBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		pterm.Fatal.Println("failed to generate pem format public key")

		return "", ""
	}

	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicBytes})

	privateBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		pterm.Fatal.Println("failed to generate pem format private key")

		return "", ""
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateBytes})

	return string(pemEncodedPub), string(pemEncoded)
}
