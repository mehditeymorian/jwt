package key

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/pterm/pterm"
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

func hmac(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)

	var base64Encoded bool
	var sizeStr string

	prompt := &survey.Confirm{Message: "generate base64 encoded key?", Default: true}

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

	hmacKey := keyGenerator.GenerateHmacKey(int(size), base64Encoded)

	pterm.Info.Println("key")
	fmt.Println(string(hmacKey))

	if saveFile {
		SaveKey("/key.txt", hmacKey)
	}

	if saveDefault {
		cfg := config.Load(configPath)
		cfg.Hmac.Key = string(hmacKey)
		cfg.Hmac.Base64Encoded = base64Encoded
		cfg.Save()
	}
}
