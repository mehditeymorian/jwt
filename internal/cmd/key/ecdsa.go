package key

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/spf13/cobra"
)

func ecdsaCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "ecdsa",
		Short:   "generate ecdsa key",
		Example: "jwt key ecdsa",
		Run:     ecdsa,
	}

	return c
}

func ecdsa(_ *cobra.Command, _ []string) {

	prompt := &survey.Select{
		Message: "select number of bits",
		Options: []string{
			"P224",
			"P256",
			"P384",
			"P521",
		},
	}

	var ellipticCurve string

	survey.AskOne(prompt, &ellipticCurve)

	publicKey, privateKey := keyGenerator.GenerateEcdsaKeys(ellipticCurve)

	dir, _ := os.Getwd()

	publicPem, err := os.Create(dir + "/public.pem")
	if err != nil {
		log.Fatalf("failed to create public pem file: %v\n", err)
	}
	defer publicPem.Close()

	publicPem.WriteString(publicKey)

	privatePem, err := os.Create(dir + "/private.pem")
	if err != nil {
		log.Fatalf("failed to create private pem file: %v\n", err)
	}
	defer privatePem.Close()

	privatePem.WriteString(privateKey)

	log.Printf(`
	Publickey: %s/public.pem
	PrivateKey: %s/private.pem
`, dir, dir)
}
