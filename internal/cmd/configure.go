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

	config.AddCommand(
		edit,
		view,
	)

	return &config
}

func view(_ *cobra.Command, _ []string) {
	config.Load().Print()
}

func edit(_ *cobra.Command, _ []string) {
	if _, err := os.Stat(config.Path); err != nil {
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
