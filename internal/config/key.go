package config

import (
	"encoding/base64"
	"regexp"

	"github.com/golang-jwt/jwt/v4"
	keyGenerator "github.com/mehditeymorian/jwt/v2/internal/key"
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

		key = keyGenerator.DecodeRsaPublicKey(temp)
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

		key = keyGenerator.DecodeEcdsaPublicKey(temp)
		err = nil
	}

	if err != nil {
		pterm.Fatal.Printf("failed to read decode key: %v\n", err)
	}

	return key
}

func (c *Config) EncodeKey(algorithm string) any {
	var err error
	var key any

	switch {
	case matchAlgorithm("(R|P)S.*", algorithm):
		pterm.Info.Println("Using RSA key for encoding")
		temp := c.Rsa.PrivateKey

		if temp == "" {
			pterm.Warning.Println("RSA key is not available, generating random rsa key")

			_, temp = keyGenerator.GenerateRsaKeys(2048)
		}

		key, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(temp))
	case matchAlgorithm("ES.*", algorithm):
		pterm.Info.Println("Using ECDSA key for encoding")

		temp := c.Ecdsa.PrivateKey

		if temp == "" {
			pterm.Warning.Println("ECDSA key is not available, generating random key")
		}

		key, err = jwt.ParseECPrivateKeyFromPEM([]byte(temp))
	case matchAlgorithm("HS.*", algorithm):
		pterm.Info.Println("Using HMAC key for encoding")

		temp := c.Hmac.Key

		if temp == "" {
			pterm.Warning.Println("HMAC key is not available, generating random key")

			temp = string(keyGenerator.GenerateHmacKey(64, true))
		}

		if c.Hmac.Base64Encoded {
			key, err = base64.StdEncoding.DecodeString(temp)
		} else {
			key = []byte(temp)
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
