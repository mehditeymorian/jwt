package key

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
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

	c.AddCommand(
		rsaCommand(),
		hmacCommand(),
		ecdsaCommand(),
	)

	return c
}

func key(_ *cobra.Command, _ []string) {

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
		rsa(nil, nil)
	case config.HMAC:
		hmac(nil, nil)
	case config.ECDSA:
		ecdsa(nil, nil)

	default:
		log.Println("this type of key is not provided")
	}

}
