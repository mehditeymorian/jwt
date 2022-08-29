package key

import (
	"log"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/spf13/cobra"
)

func hmacCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "hmac",
		Short:   "generate hmac key",
		Example: "jwt key hmac",
		Run:     hmac,
	}

	return c
}

func hmac(_ *cobra.Command, _ []string) {

	var base64Encoded bool
	var sizeStr string

	prompt := &survey.Confirm{Message: "generate base64 encoded key?"}

	survey.AskOne(prompt, &base64Encoded)

	promptSize := &survey.Select{
		Message: "select key size",
		Options: []string{
			"64",
			"128",
			"256",
			"512",
			"1024",
		},
	}

	survey.AskOne(promptSize, &sizeStr)

	size, _ := strconv.ParseInt(sizeStr, 10, 64)

	dir, _ := os.Getwd()

	hmacKey := keyGenerator.GenerateHmacKey(int(size), base64Encoded)

	keyFile, err := os.Create(dir + "/key.txt")
	if err != nil {
		log.Fatalf("failed to create key file: %v\n", err)
	}
	defer keyFile.Close()

	keyFile.Write(hmacKey)

	log.Printf(`
	Publickey: %s/key.txt
`, dir)
}
