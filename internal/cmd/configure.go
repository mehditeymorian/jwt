package cmd

import (
	"os"
	"os/exec"

	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Configure() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "config jwt cli",
		Long:  "config jwt cli",
		Run:   configure,
	}
}

func configure(_ *cobra.Command, _ []string) {
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
