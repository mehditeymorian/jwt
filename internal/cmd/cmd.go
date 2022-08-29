package cmd

import "github.com/spf13/cobra"

func SetConfigFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("config", "c", "", "jwt configuration")
}

func GetConfigPath(cmd *cobra.Command) string {
	path, _ := cmd.Flags().GetString("config")

	return path
}
