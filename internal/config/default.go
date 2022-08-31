package config

import (
	"github.com/mehditeymorian/jwt/internal/key"
)

func Default() Config {
	rsaPublic, rsaPrivate := key.GenerateRsaKeys(1024)
	hmacKey := key.GenerateHmacKey(256, true)
	ecPublic, ecPrivate := key.GenerateEcdsaKeys("P256")

	return Config{
		Algorithms: []string{
			"HS256",
			"HS384",
			"HS512",
			"RS256",
			"RS384",
			"RS512",
			"ES256",
			"ES384",
			"ES512",
			"PS256",
			"PS384",
			"PS512",
		},
		Expirations: []string{
			"5m",
			"10m",
			"1h",
			"24h",
			"720h",
			"8760h",
		},

		Interactive: true,

		SigningMethod: RSA,
		Algorithm:     "RS512",

		Rsa: &Rsa{
			PublicKey:  rsaPublic,
			PrivateKey: rsaPrivate,
		},
		Hmac: &HMac{
			Key:           string(hmacKey),
			Base64Encoded: true,
		},
		Ecdsa: &Ecdsa{
			PublicKey:  ecPublic,
			PrivateKey: ecPrivate,
		},
	}
}
