package key

import (
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
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

func rsa(c *cobra.Command, _ []string) {
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)

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

	if saveFile {
		SaveKey("/public.pem", []byte(publicKey))
		SaveKey("/private.pem", []byte(privateKey))
	}

	if saveDefault {

	}

}
