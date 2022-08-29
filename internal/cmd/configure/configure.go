package configure

import (
	"os"
	"os/exec"

	"github.com/mehditeymorian/jwt/internal/cmd"
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

	edit := &cobra.Command{
		Use:   "edit",
		Short: "edit jwt config",
		Run:   edit,
	}

	view := &cobra.Command{
		Use:   "view",
		Short: "edit jwt config",
		Run:   view,
	}

	c.AddCommand(
		edit,
		view,
	)

	return c
}

func view(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)

	config.Load(configPath).Print()
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
