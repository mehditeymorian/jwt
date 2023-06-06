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
		Short:   "open jwt default config in vim",
		Example: "jwt config edit",
		Run:     edit,
	}

	return c
}

func edit(c *cobra.Command, _ []string) {
	openVim(c)
}

func openVim(c *cobra.Command) {
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
