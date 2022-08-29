package config

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/golang-jwt/jwt/v4"
)

func (c Config) DecodeKey(algorithm string) any {
	var err error
	var key any

	switch {
	case matchAlgorithm("(R|P)S.*", algorithm):
		key, err = jwt.ParseRSAPublicKeyFromPEM([]byte(c.Rsa.PublicKey))
	case matchAlgorithm("HS.*", algorithm):
		if c.Hmac.Base64Encoded {
			key, err = base64.StdEncoding.DecodeString(c.Hmac.Key)
		} else {
			key = []byte(c.Hmac.Key)
		}
	case matchAlgorithm("ES.*", algorithm):
		key, err = jwt.ParseECPublicKeyFromPEM([]byte(c.Ecdsa.PublicKey))
	}

	if err != nil {
		panic(fmt.Errorf("failed to read decode key from config: %w", err))
	}

	return key
}

func (c Config) EncodeKey() any {
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
		panic(fmt.Errorf("failed to read encode key from config: %w", err))
	}

	return key
}

func matchAlgorithm(pattern string, algorithm string) bool {
	matched, _ := regexp.MatchString(pattern, algorithm)

	return matched
}
