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

func rsaCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "rsa",
		Short:   "generate rsa key",
		Example: "jwt key rsa",
		Run:     rsa,
	}
	c.Flags().IntP("bits", "b", 1024, "number of bits")

	return c
}

func rsa(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)
	cfg := config.Load(configPath)

	var bits int

	if cfg.Interactive {
		bits = askRsaOptions()
	} else {
		bits = flagRsaOptions(c)
	}

	cfg.PrintMode()
	pterm.Info.Println("bits: " + pterm.Blue(bits))

	publicKey, privateKey := keyGenerator.GenerateRsaKeys(bits)

	publicBox := pterm.DefaultBox.WithTitle("Public Key").Sprint(publicKey)
	privateBox := pterm.DefaultBox.WithTitle("Private Key").Sprint(privateKey)
	render, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{{{Data: publicBox}, {Data: privateBox}}}).Srender()
	pterm.Println(render)

	if saveFile {
		SaveKey("/public.pem", []byte(publicKey))
		SaveKey("/private.pem", []byte(privateKey))
	}

	if saveDefault {
		cfg.Rsa.PublicKey = publicKey
		cfg.Rsa.PrivateKey = privateKey
		cfg.Save()
	}

}

func askRsaOptions() int {
	var bitsStr string

	prompt := &survey.Select{
		Message: "select number of bits",
		Options: []string{
			"512",
			"1024",
			"2048",
			"4096",
		},
	}

	survey.AskOne(prompt, &bitsStr)

	bits, _ := strconv.ParseInt(bitsStr, 10, 64)

	return int(bits)
}

func flagRsaOptions(c *cobra.Command) int {
	bits, _ := c.Flags().GetInt("bits")

	return bits
}
