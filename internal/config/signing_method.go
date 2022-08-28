package config

type SigningMethod string

const (
	RSA  SigningMethod = "rsa"
	HMAC SigningMethod = "hmac"
)

type Rsa struct {
	PublicKey  string `koanf:"public_key"`
	PrivateKey string `koanf:"private_key"`
}

type HMac struct {
	Key           string `koanf:"key"`
	Base64Encoded bool   `koanf:"base64_encoded"`
}
