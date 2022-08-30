package edit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func methodCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "method",
		Short:   "config signing method",
		Example: "jwt config edit method",
		Run:     method,
	}

	return c
}

func method(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	prompt := &survey.Select{
		Message: "select signing method",
		Options: []string{
			string(config.RSA),
			string(config.ECDSA),
			string(config.HMAC),
		},
	}

	var result string

	survey.AskOne(prompt, &result)

	cfg.SigningMethod = config.SigningMethod(result)

	cfg.Save()
}
