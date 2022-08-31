package encode

import (
	"log"

	"github.com/mehditeymorian/jwt/internal/cmd"
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
	root.Flags().StringP("expiration", "e", "1h", "token expires after how long")
	root.Flags().StringP("subject", "s", "", "token subject")
	root.Flags().StringP("issuer", "i", "", "token issuer")
	root.Flags().StringP("audience", "a", "", "token audience")
	root.Flags().StringP("algorithm", "A", "", "signing method algorithm")

	return root
}

func main(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	var encode model.Encode

	if cfg.Interactive {
		encode = askOptions(cfg)
	} else {
		encode = flagOptions(c)
	}

	token, err := jwt.Encode(encode, cfg.EncodeKey())
	if err != nil {
		log.Fatalf("failed to generate JWT token: %v", err)
	}

	log.Printf("Token: %s\n", token)
}

func askOptions(cfg *config.Config) model.Encode {
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
		{
			Name: "Algorithm",
			Prompt: &survey.Select{
				Message: "Algorithm",
				Options: cfg.AlgorithmForMethod(),
			},
		},
	}

	survey.Ask(qs, &encode, nil)

	return encode
}

func flagOptions(c *cobra.Command) model.Encode {
	exp, _ := c.Flags().GetString("expiration")
	sub, _ := c.Flags().GetString("subject")
	iss, _ := c.Flags().GetString("issuer")
	aud, _ := c.Flags().GetString("audience")
	alg, _ := c.Flags().GetString("algorithm")

	return model.Encode{
		Algorithm:  alg,
		Expiration: exp,
		Issuer:     iss,
		Subject:    sub,
		Audience:   aud,
	}
}
