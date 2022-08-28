package config

import (
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func (c Config) DecodeKey() any {
	var err error
	var key any

	switch c.SigningMethod {
	case RSA:
		key, err = jwt.ParseRSAPublicKeyFromPEM([]byte(c.Rsa.PublicKey))
	case HMAC:
		if c.Hmac.Base64Encoded {
			key, err = base64.StdEncoding.DecodeString(c.Hmac.Key)
		} else {
			key = []byte(c.Hmac.Key)
		}
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
	default:
		key = c.DecodeKey()
	}

	if err != nil {
		panic(fmt.Errorf("failed to read encode key from config: %w", err))
	}

	return key
}
