package key

import (
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
	c.Flags().BoolP("base64", "b", true, "produce base64 encoded key")
	c.Flags().IntP("size", "s", 64, "size of key")

	return c
}

func hmac(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)
	cfg := config.Load(configPath)

	var size int
	var base64Encoded bool

	if cfg.Interactive {
		size, base64Encoded = askHmacOptions()
	} else {
		size, base64Encoded = flagHmacOptions(c)
	}

	cfg.PrintMode()
	pterm.Info.Println("size: " + pterm.Blue(size))
	pterm.Info.Println("base64Encoded: " + pterm.Blue(base64Encoded))

	hmacKey := keyGenerator.GenerateHmacKey(size, base64Encoded)

	pterm.Info.Println("key: " + pterm.Blue(string(hmacKey)))

	if saveFile {
		SaveKey("/key.txt", hmacKey)
	}

	if saveDefault {
		cfg.Hmac.Key = string(hmacKey)
		cfg.Hmac.Base64Encoded = base64Encoded
		cfg.Save()
	}
}

func askHmacOptions() (int, bool) {
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

	return int(size), base64Encoded
}

func flagHmacOptions(c *cobra.Command) (int, bool) {
	size, _ := c.Flags().GetInt("size")
	base64, _ := c.Flags().GetBool("base64")

	return size, base64
}
