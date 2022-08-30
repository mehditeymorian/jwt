package key

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/spf13/cobra"
)

func Key() *cobra.Command {
	c := &cobra.Command{
		Use:     "key",
		Short:   "generate different jwt keys",
		Run:     key,
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

func key(cmd *cobra.Command, args []string) {

	prompt := &survey.Select{
		Message: "select key type",
		Options: []string{
			string(config.RSA),
			string(config.HMAC),
			string(config.ECDSA),
		},
	}

	var selected string

	survey.AskOne(prompt, &selected)

	switch config.SigningMethod(selected) {
	case config.RSA:
		rsa(cmd, args)
	case config.HMAC:
		hmac(cmd, args)
	case config.ECDSA:
		ecdsa(cmd, args)

	default:
		log.Println("this type of key is not provided")
	}

}

func SaveKey(filename string, content []byte) {
	dir, _ := os.Getwd()

	file, err := os.Create(dir + filename)
	if err != nil {
		log.Fatalf("failed to save key in file: %v\n", err)
	}
	defer file.Close()

	file.Write(content)

	log.Printf("key stored in %s/public.pem\n", dir)
}
