package cmd

import "github.com/spf13/cobra"

func SetConfigFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("config", "c", "", "jwt configuration")
}

func GetConfigPath(cmd *cobra.Command) string {
	path, _ := cmd.Flags().GetString("config")

	return path
}

func SetKeyFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP("file", "f", false, "save keys as file")
	cmd.PersistentFlags().BoolP("default", "d", false, "save keys in default config")
}

func GetKeySaveOptions(cmd *cobra.Command) (saveFile bool, saveDefault bool) {
	saveFile, _ = cmd.Flags().GetBool("file")
	saveDefault, _ = cmd.Flags().GetBool("default")

	return saveFile, saveDefault
}
