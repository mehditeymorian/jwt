package config

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"regexp"

	"github.com/golang-jwt/jwt/v4"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/pterm/pterm"
)

func (c *Config) DecodeKey(algorithm string) any {
	var err error
	var key any

	switch {
	case matchAlgorithm("(R|P)S.*", algorithm):
		pterm.Info.Println("Using RSA key for decoding")

		temp := c.Rsa.PublicKey
		if temp == "" {
			pterm.Warning.Println("RSA key is not available, generating random rsa key")

			temp, _ = keyGenerator.GenerateRsaKeys(2048)
		}

		block, _ := pem.Decode([]byte(temp))
		if block == nil || block.Type != "PUBLIC KEY" {
			pterm.Fatal.Println("failed to decode pem value containing public key")
		}

		var publicKey rsa.PublicKey
		if _, err := asn1.Unmarshal(block.Bytes, &publicKey); err != nil {
			pterm.Fatal.Printf("failed to unmarshal asn1 data to public key: %v\n", err)
		}

		key = &publicKey
		err = nil
	case matchAlgorithm("HS.*", algorithm):
		pterm.Info.Println("Using HMAC key for decoding")

		temp := c.Hmac.Key
		if temp == "" {
			pterm.Warning.Println("Hmac key is not available, generating random hmac key")

			temp = string(keyGenerator.GenerateHmacKey(64, true))
		}

		if c.Hmac.Base64Encoded {
			key, err = base64.StdEncoding.DecodeString(temp)
		} else {
			key = []byte(temp)
		}
	case matchAlgorithm("ES.*", algorithm):
		pterm.Info.Println("Using ECDSA key for decoding")

		temp := c.Ecdsa.PublicKey
		if temp == "" {
			pterm.Warning.Println("Ecdsa key is not available, generating random Ecdsa key")

			temp, _ = keyGenerator.GenerateEcdsaKeys("P256")
		}

		block, _ := pem.Decode([]byte(temp))
		if block == nil || block.Type != "PUBLIC KEY" {
			pterm.Fatal.Println("failed to decode pem value containing public key")
		}

		parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			pterm.Fatal.Printf("failed to unmarshal pkix data to public key: %v\n", err)
		}

		key = parsedKey.(*ecdsa.PublicKey)
		err = nil
	}

	if err != nil {
		pterm.Fatal.Printf("failed to read decode key: %v\n", err)
	}

	return key
}

func (c *Config) EncodeKey() any {
	var err error
	var key any

	switch c.SigningMethod {
	case RSA:
		key, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(c.Rsa.PrivateKey))
	case ECDSA:
		key, err = jwt.ParseECPrivateKeyFromPEM([]byte(c.Ecdsa.PrivateKey))
	case HMAC:
		if c.Hmac.Base64Encoded {
			key, err = base64.StdEncoding.DecodeString(c.Hmac.Key)
		} else {
			key = []byte(c.Hmac.Key)
		}
	}

	if err != nil {
		pterm.Fatal.Printf("failed to read encode key from config: %v\n", err)
	}

	return key
}

func matchAlgorithm(pattern string, algorithm string) bool {
	matched, _ := regexp.MatchString(pattern, algorithm)

	return matched
}
