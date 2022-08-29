package cmd

import (
	"log"

	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/mehditeymorian/jwt/internal/jwt"
	"github.com/mehditeymorian/jwt/internal/model"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func Encode() *cobra.Command {
	root := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "encode",
		Short: "Create JWT Token",
		Long:  "Create JWT Token",
		Run:   main,
	}
	SetConfigFlag(root)

	return root
}

func main(cmd *cobra.Command, _ []string) {
	configPath := GetConfigPath(cmd)

	cfg := config.Load(configPath)

	var encode model.Encode

	qs := []*survey.Question{
		{
			Name: "Expiration",
			Prompt: &survey.Select{ //nolint:exhaustruct,exhaustivestruct
				Message: "Expiration",
				Options: cfg.Expirations,
				Default: cfg.Expirations[0],
			},
		},
		{
			Name: "Subject",
			Prompt: &survey.Input{ //nolint:exhaustruct,exhaustivestruct
				Message: "Subject",
			},
		},
		{
			Name: "Issuer",
			Prompt: &survey.Input{ //nolint:exhaustruct,exhaustivestruct
				Message: "Issuer",
			},
		},
		{
			Name: "Audience",
			Prompt: &survey.Input{ //nolint:exhaustruct,exhaustivestruct
				Message: "Audience",
			},
		},
	}

	survey.Ask(qs, &encode, nil)

	encode.Algorithm = cfg.Algorithm

	token, err := jwt.Encode(encode, cfg.EncodeKey())
	if err != nil {
		log.Fatalf("failed to generate JWT token: %v", err)
	}

	log.Printf("Token: %s\n", token)
}
