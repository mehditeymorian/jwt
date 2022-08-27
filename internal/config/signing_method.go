package config

type SigningMethod string

const (
	RSA  SigningMethod = "rsa"
	HMac SigningMethod = "hmac"
)

type Rsa struct {
	PublicKey  string
	PrivateKey string
}
