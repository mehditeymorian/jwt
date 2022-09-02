package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"

	"github.com/pterm/pterm"
)

func GenerateRsaKeys(bits int) (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		pterm.Fatal.Println("failed to generate rsa keys")

		return "", ""
	}

	publicKey, err := asn1.Marshal(privateKey.PublicKey)
	if err != nil {
		pterm.Fatal.Printf("failed to generate public key: %v\n", err)
	}

	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKey})

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	return string(pemEncodedPub), string(pemEncoded)
}

func DecodeRsaPublicKey(public string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(public))
	if block == nil || block.Type != "PUBLIC KEY" {
		pterm.Fatal.Println("failed to decode pem value containing public key")
	}

	var publicKey rsa.PublicKey
	if _, err := asn1.Unmarshal(block.Bytes, &publicKey); err != nil {
		pterm.Fatal.Printf("failed to unmarshal asn1 data to public key: %v\n", err)
	}

	return &publicKey
}
