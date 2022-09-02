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
		ellipticCurve = elliptic.P256()
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

func DecodeEcdsaPublicKey(public string) *ecdsa.PublicKey {
	block, _ := pem.Decode([]byte(public))
	if block == nil || block.Type != "PUBLIC KEY" {
		pterm.Fatal.Println("failed to decode pem value containing public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pterm.Fatal.Printf("failed to unmarshal pkix data to public key: %v\n", err)
	}

	return parsedKey.(*ecdsa.PublicKey)
}
