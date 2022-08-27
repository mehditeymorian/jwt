package jwt

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mehditeymorian/jwt/internal/model"
)

func Encode(encode model.Encode) (string, error) {
	publicKeyFile, err := os.Open(encode.PublicKeyPath)
	if err != nil {
		return "", fmt.Errorf("failed to open public key file: %w", err)
	}

	signingKey, err := ioutil.ReadAll(publicKeyFile)
	if err != nil {
		return "", fmt.Errorf("failed to read public key: %w", err)
	}

	exp, _ := time.ParseDuration(encode.Expiration)

	claims := jwt.MapClaims{
		"iss":  encode.Issuer,
		"exp":  time.Now().Add(exp).Unix(),
		"data": encode.Payload,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(encode.Algorithm), claims)

	signedString, err := token.SignedString(signingKey)

	return signedString, fmt.Errorf("failed to sign claims with key: %w", err)
}
