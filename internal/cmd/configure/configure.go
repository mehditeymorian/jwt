package configure

import (
	"github.com/mehditeymorian/jwt/internal/cmd/configure/edit"
	"github.com/mehditeymorian/jwt/internal/cmd/configure/view"
	"github.com/spf13/cobra"
)

func Configure() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "config jwt cli",
		Long:  "config jwt cli",
	}

	c.AddCommand(
		edit.Command(),
		view.Command(),
	)

	return c
}
