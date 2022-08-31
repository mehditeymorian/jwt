package key

import (
	"os"

	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "key",
		Short:   "generate different jwt keys",
		Example: "jwt key rsa|hmac|ecdsa",
	}
	cmd.SetKeyFlags(c)

	c.AddCommand(
		rsaCommand(),
		hmacCommand(),
		ecdsaCommand(),
	)

	return c
}

func SaveKey(filename string, content []byte) {
	dir, _ := os.Getwd()

	file, err := os.Create(dir + filename)
	if err != nil {
		pterm.Fatal.Printf("failed to save key in file: %v\n", err)
	}
	defer file.Close()

	file.Write(content)

	pterm.Success.Printf("key stored in %s/public.pem\n", dir)
}
