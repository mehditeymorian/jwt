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
	c.Flags().StringP("curve", "c", "P256", "elliptic curve")

	return c
}

func ecdsa(c *cobra.Command, _ []string) {
	configPath := cmd.GetConfigPath(c)
	saveFile, saveDefault := cmd.GetKeySaveOptions(c)
	cfg := config.Load(configPath)

	var ellipticCurve string
	if cfg.Interactive {
		ellipticCurve = askEcdsaOptions()
	} else {
		ellipticCurve = flagEcdsaOptions(c)
	}

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
		cfg.Ecdsa.PublicKey = publicKey
		cfg.Ecdsa.PrivateKey = privateKey
		cfg.Save()
	}
}

func askEcdsaOptions() string {
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

	return ellipticCurve
}

func flagEcdsaOptions(c *cobra.Command) string {
	curve, _ := c.Flags().GetString("curve")

	return curve
}
