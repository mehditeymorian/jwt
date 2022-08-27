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

	return root
}

func main(cmd *cobra.Command, args []string) {
	cfg := config.Load()

	var encode model.Encode

	qs := []*survey.Question{
		{
			Name: "Algorithm",
			Prompt: &survey.Select{ //nolint:exhaustruct,exhaustivestruct
				Message: "Algorithm",
				Options: cfg.JWT.Algorithms,
				Default: cfg.JWT.Algorithms[0],
			},
		},
		{
			Name: "Expiration",
			Prompt: &survey.Select{ //nolint:exhaustruct,exhaustivestruct
				Message: "Expiration",
				Options: cfg.JWT.Expirations,
				Default: cfg.JWT.Expirations[0],
			},
		},
		{
			Name: "PayloadStr",
			Prompt: &survey.Multiline{ //nolint:exhaustruct,exhaustivestruct
				Message: "Enter Payload Fields",
				Help:    "Format: KEY:VALUE",
			},
		},
		{
			Name: "Issuer",
			Prompt: &survey.Input{ //nolint:exhaustruct,exhaustivestruct
				Message: "Issuer",
			},
		},
		{
			Name: "PublicKeyPath",
			Prompt: &survey.Input{ //nolint:exhaustruct,exhaustivestruct
				Message: "Public Key File Path",
			},
		},
	}

	survey.Ask(qs, &encode, nil)

	encode.Execute()

	token, err := jwt.Encode(encode)
	if err != nil {
		log.Fatalf("failed to generate JWT token: %w", err)
	}

	log.Printf("Token: %s\n", token)
}
