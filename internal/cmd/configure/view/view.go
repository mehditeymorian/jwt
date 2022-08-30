package view

import (
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "view",
		Short: "edit jwt config",
		Run:   view,
	}

	return c
}

func view(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	config.Load(configPath).Print()
}
