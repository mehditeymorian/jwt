package set

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "set",
		Short:   "set config",
		Example: "jwt set interactive true",
		Run:     set,
	}

	return c
}

func set(cmd *cobra.Command, args []string) {

}
