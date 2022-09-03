package set

import (
	"strconv"
	"strings"

	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "set",
		Short:   "set config",
		Example: "jwt set interactive true",
		Run:     set,
	}

	return c
}

func set(c *cobra.Command, args []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	if len(args) < 2 {
		pterm.Fatal.Println("not enough argument")
	}

	keys := strings.Split(args[0], ".")

	switch keys[0] {
	case "interactive":
		setInteractive(cfg, args)
	}

	cfg.Save()
}

func setInteractive(cfg *config.Config, args []string) {
	interactive, err := strconv.ParseBool(args[1])
	if err != nil {
		pterm.Fatal.Println("failed to parse interactive value")
	}

	cfg.Interactive = interactive
}
