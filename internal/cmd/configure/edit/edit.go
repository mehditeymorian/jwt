package edit

import (
	"os"
	"os/exec"

	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "edit",
		Short:   "edit jwt config",
		Example: "jwt config edit",
		Run:     edit,
	}

	c.AddCommand(
		algorithmCommand(),
		methodCommand())

	return c
}

func edit(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	if _, err := os.Stat(configPath); err != nil {
		cmd := exec.Command("sudo", "mkdir", "-p", config.Dir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()

		cmd = exec.Command("sudo", "touch", config.Path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	cmd := exec.Command("vim", config.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
