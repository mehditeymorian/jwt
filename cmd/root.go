/*
Package cmd
Copyright © 2022 Mehdi Teymorian
*/
package cmd

import (
	"os"

	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/cmd/config"
	"github.com/mehditeymorian/jwt/internal/cmd/decode"
	"github.com/mehditeymorian/jwt/internal/cmd/encode"
	"github.com/mehditeymorian/jwt/internal/cmd/gen"
	"github.com/mehditeymorian/jwt/internal/cmd/set"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{ //nolint:exhaustruct,exhaustivestruct
		Use:   "jwt",
		Short: "JWT Encoder and Decoder",
		Long:  `Command and Command JWT Tokens`,
	}
	cmd.SetConfigFlag(rootCmd)

	rootCmd.AddCommand(
		encode.Command(),
		decode.Command(),
		config.Command(),
		gen.Command(),
		set.Command(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
