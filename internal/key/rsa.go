package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRsaKeys(bits int) (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", ""
	}

	publicKey := &privateKey.PublicKey
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKey)})

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	return string(pemEncodedPub), string(pemEncoded)
}
