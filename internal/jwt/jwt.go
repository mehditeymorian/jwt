package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mehditeymorian/jwt/v2/internal/config"
	"github.com/mehditeymorian/jwt/v2/internal/model"
)

func Encode(encode model.Encode, key any) (string, error) {
	exp, _ := time.ParseDuration(encode.Expiration)

	claims := jwt.MapClaims{
		"iss": encode.Issuer,
		"exp": time.Now().Add(exp).Unix(),
		"iat": time.Now().Unix(),
		"sub": encode.Subject,
		"aud": encode.Audience,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(encode.Algorithm), claims)

	signedString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign claims with key: %w", err)
	}

	return signedString, nil
}

func Decode(strToken string, cfg *config.Config) (*jwt.Token, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {

		return cfg.DecodeKey(token.Method.Alg()), nil
	})

	return token, err
}
