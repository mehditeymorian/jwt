package config

import (
	"encoding/base64"
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
		temp := c.Rsa.PublicKey
		if temp == "" {
			pterm.Warning.Println("RSA key is not available, generating random rsa key")

			temp, _ = keyGenerator.GenerateRsaKeys(2048)
		}

		key, err = jwt.ParseRSAPublicKeyFromPEM([]byte(temp))
	case matchAlgorithm("HS.*", algorithm):
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
		temp := c.Ecdsa.PublicKey
		if temp == "" {
			pterm.Warning.Println("Ecdsa key is not available, generating random Ecdsa key")

			key, _ = keyGenerator.GenerateEcdsaKeys("P256")
		}

		key, err = jwt.ParseECPublicKeyFromPEM([]byte(temp))
	}

	if err != nil {
		pterm.Fatal.Printf("failed to read decode key from config: %v\n", err)
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
