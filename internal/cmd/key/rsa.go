package key

import (
	"log"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/spf13/cobra"
)

func rsaCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "rsa",
		Short:   "generate rsa key",
		Example: "jwt key rsa",
		Run:     rsa,
	}

	return c
}

func rsa(_ *cobra.Command, _ []string) {

	prompt := &survey.Select{
		Message: "select number of bits",
		Options: []string{
			"512",
			"1024",
			"2048",
			"4096",
		},
	}

	var bitsStr string

	survey.AskOne(prompt, &bitsStr)

	bits, _ := strconv.ParseInt(bitsStr, 10, 64)

	publicKey, privateKey := keyGenerator.GenerateRsaKeys(int(bits))

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
