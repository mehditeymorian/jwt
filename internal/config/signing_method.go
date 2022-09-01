package config

type SigningMethod string

const (
	RSA   SigningMethod = "rsa"
	HMAC  SigningMethod = "hmac"
	ECDSA SigningMethod = "ecdsa"
)

type Rsa struct {
	PublicKey  string `koanf:"public_key" yaml:"public_key"`
	PrivateKey string `koanf:"private_key" yaml:"private_key"`
}

type HMac struct {
	Key           string `koanf:"key" yaml:"key"`
	Base64Encoded bool   `koanf:"base64_encoded" yaml:"base64_encoded"`
}

type Ecdsa struct {
	PublicKey  string `koanf:"public_key" yaml:"public_key"`
	PrivateKey string `koanf:"private_key" yaml:"private_key"`
}
