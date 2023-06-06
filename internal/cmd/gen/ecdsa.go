package gen

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/v2/internal/cmd"
	"github.com/mehditeymorian/jwt/v2/internal/config"
	keyGenerator "github.com/mehditeymorian/jwt/v2/internal/key"
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
	c.Flags().StringP("elliptic", "e", "P256", "elliptic curve")

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

	cfg.PrintMode()
	pterm.Println("elliptic curve: " + pterm.Blue(ellipticCurve))

	publicKey, privateKey := keyGenerator.GenerateEcdsaKeys(ellipticCurve)

	pterm.Info.Println("Public Key")
	fmt.Println(publicKey)
	pterm.Info.Println("Private Key")
	fmt.Println(privateKey)

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
		Message: "select elliptic curve",
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
	curve, _ := c.Flags().GetString("elliptic")

	return curve
}
