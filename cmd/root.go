/*Package cmd
Copyright © 2022 Mehdi Teymorian

*/
package cmd

import (
	"os"

	"jwt-cli/internal/cmd"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{ //nolint:exhaustruct,exhaustivestruct
		Use:   "jwt-cli",
		Short: "JWT Encoder and Decoder",
		Long:  `Encode and Decode JWT Tokens`,
	}

	rootCmd.AddCommand(
		cmd.Encode(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
