package config

import (
	"github.com/mehditeymorian/jwt/v2/internal/cmd"
	"github.com/mehditeymorian/jwt/v2/internal/cmd/config/edit"
	"github.com/mehditeymorian/jwt/v2/internal/cmd/config/set"
	"github.com/mehditeymorian/jwt/v2/internal/config"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "config",
		Short:   "view and edit config",
		Example: "jwt config",
		Run:     view,
	}

	c.AddCommand(
		edit.Command(),
		set.Command(),
	)

	return c
}

func view(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	config.Load(configPath).Print()
}
