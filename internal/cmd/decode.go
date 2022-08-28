package cmd

import (
	"encoding/json"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/mehditeymorian/jwt/internal/jwt"
	"github.com/spf13/cobra"
)

func Decode() *cobra.Command {
	return &cobra.Command{
		Use:   "decode",
		Short: "decode jwt token",
		Long:  "decode jwt token",
		Run:   decode,
	}
}

func decode(_ *cobra.Command, _ []string) {
	cfg := config.Load()

	strToken := ""

	prompt := &survey.Input{Message: "JWT Token"}

	survey.AskOne(prompt, &strToken)

	token, err := jwt.Decode(strToken, cfg.DecodeKey(), cfg.Algorithm)
	if err != nil {
		log.Fatalf("failed to decode token: %v", err)
	}

	headers, err := json.MarshalIndent(token.Header, "", "\t")
	if err != nil {
		log.Fatalf("failed to decode token: %v", err)
	}

	payload, err := json.MarshalIndent(token.Claims, "", "\t")
	if err != nil {
		log.Fatalf("failed to decode token: %v", err)
	}

	if token.Valid {
		log.Println("Token is Valid.")
	} else {
		log.Println("Token is Invalid.")
	}

	log.Printf("headers: %s,\npayload: %s\n", headers, payload)
}
