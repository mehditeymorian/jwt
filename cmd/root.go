/*
Package cmd
Copyright © 2022 Mehdi Teymorian
*/
package cmd

import (
	"os"

	"github.com/mehditeymorian/jwt/internal/cmd"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{ //nolint:exhaustruct,exhaustivestruct
		Use:   "jwt",
		Short: "JWT Encoder and Decoder",
		Long:  `Encode and Decode JWT Tokens`,
	}

	rootCmd.AddCommand(
		cmd.Encode(),
		cmd.Decode(),
		cmd.Configure(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
