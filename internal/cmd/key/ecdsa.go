package key

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	keyGenerator "github.com/mehditeymorian/jwt/internal/key"
	"github.com/pterm/pterm"
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

func ecdsa(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)

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

	publicBox := pterm.DefaultBox.WithTitle("Public Key").Sprint(publicKey)
	privateBox := pterm.DefaultBox.WithTitle("Private Key").Sprint(privateKey)
	render, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{{{Data: publicBox}, {Data: privateBox}}}).Srender()
	pterm.Println(render)

	if saveFile {
		SaveKey("/public.pem", []byte(publicKey))
		SaveKey("/private.pem", []byte(privateKey))
	}

	if saveDefault {
		cfg := config.Load(configPath)
		cfg.Ecdsa.PublicKey = publicKey
		cfg.Ecdsa.PrivateKey = privateKey
		cfg.Save()
	}
}
