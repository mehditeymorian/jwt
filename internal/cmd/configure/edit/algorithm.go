package edit

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func algorithmCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "algo",
		Short:   "set encode algorithm",
		Example: "jwt config edit algo",
		Run:     algorithm,
	}

	return c
}

func algorithm(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	prompt := &survey.Select{
		Message: "select algorithm",
		Options: cfg.Algorithms,
	}

	var alg string

	survey.AskOne(prompt, &alg)

	cfg.Algorithm = alg

	cfg.Save()
}
