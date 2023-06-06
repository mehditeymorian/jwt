package encode

import (
	"github.com/mehditeymorian/jwt/v2/internal/cmd"
	"github.com/mehditeymorian/jwt/v2/internal/config"
	"github.com/mehditeymorian/jwt/v2/internal/jwt"
	"github.com/mehditeymorian/jwt/v2/internal/model"
	"github.com/pterm/pterm"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	root := &cobra.Command{ //nolint:exhaustivestruct
		Use:     "encode",
		Short:   "encode standards claims into jwt token",
		Example: `jwt encode -e 2h -s finance -i "jwt-cli" -a people -A RS512`,
		Run:     main,
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

	token, err := jwt.Encode(encode, cfg.EncodeKey(encode.Algorithm))
	if err != nil {
		pterm.Fatal.Printf("failed to generate JWT token: %v\n", err)
	}

	pterm.Success.Println("token: " + pterm.Blue(token))
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
				Options: cfg.Algorithms,
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
	if alg == "" {
		pterm.Fatal.Println("you have to specify algorithm")
	}

	return model.Encode{
		Algorithm:  alg,
		Expiration: exp,
		Issuer:     iss,
		Subject:    sub,
		Audience:   aud,
	}
}
