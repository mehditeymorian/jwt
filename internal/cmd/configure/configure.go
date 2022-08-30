package configure

import (
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/cmd/configure/edit"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Configure() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "config jwt cli",
		Long:  "config jwt cli",
		Run:   view,
	}

	view := &cobra.Command{
		Use:   "view",
		Short: "edit jwt config",
		Run:   view,
	}

	c.AddCommand(
		edit.Command(),
		view,
	)

	return c
}

func view(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	config.Load(configPath).Print()
}
