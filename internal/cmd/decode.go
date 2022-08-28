package cmd

import (
	"encoding/json"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/mehditeymorian/jwt/internal/jwt"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
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

	if token.Valid {
		log.Println("Token is Valid.")
	} else {
		log.Println("Token is Invalid.")
	}

	result := make(map[string]any)
	result["headers"] = token.Header
	result["payload"] = token.Claims

	indent, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling result: %v", err)
	}

	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Decoded JWT Token ================
	%s
	===================================================
	`
	log.Printf(cfgStrTemplate, string(indent))
}
