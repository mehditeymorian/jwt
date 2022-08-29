package cmd

import (
	"os"
	"os/exec"

	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Configure() *cobra.Command {
	config := cobra.Command{
		Use:   "config",
		Short: "config jwt cli",
		Long:  "config jwt cli",
		Run:   view,
	}

	config.Flags().StringP("config", "c", "", "jwt configuration")

	edit := &cobra.Command{
		Use:   "edit",
		Short: "edit jwt config",
		Run:   edit,
	}

	edit.Flags().StringP("config", "c", "", "jwt configuration")

	view := &cobra.Command{
		Use:   "view",
		Short: "edit jwt config",
		Run:   view,
	}

	view.Flags().StringP("config", "c", "", "jwt configuration")

	config.AddCommand(
		edit,
		view,
	)

	return &config
}

func view(c *cobra.Command, _ []string) {
	configPath, _ := c.Flags().GetString("config")

	config.Load(configPath).Print()
}

func edit(c *cobra.Command, _ []string) {
	configPath, _ := c.Flags().GetString("config")

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
