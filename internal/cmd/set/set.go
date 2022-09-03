package set

import (
	"strconv"
	"strings"

	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "set",
		Short:   "set config",
		Example: "jwt set interactive true",
		Run:     set,
	}

	return c
}

func set(c *cobra.Command, args []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	if len(args) < 2 {
		pterm.Fatal.Println("not enough argument")
	}

	keys := strings.Split(args[0], ".")

	switch keys[0] {
	case "interactive":
		setInteractive(cfg, args)
	case "rsa":
		if len(keys) < 2 {
			pterm.Fatal.Printf("%s field is not specified\n", keys[0])
		}
		setRsa(cfg, keys[1], args)
	case "hmac":
		if len(keys) < 2 {
			pterm.Fatal.Printf("%s field is not specified\n", keys[0])
		}
		setHmac(cfg, keys[1], args)
	case "ecdsa":
		if len(keys) < 2 {
			pterm.Fatal.Printf("%s field is not specified\n", keys[0])
		}
		setEcdsa(cfg, keys[1], args)
	default:
		pterm.Warning.Println("field %s not found")
	}

	cfg.Save()
}

func setInteractive(cfg *config.Config, args []string) {
	interactive, err := strconv.ParseBool(args[1])
	if err != nil {
		pterm.Fatal.Println("failed to parse interactive value")
	}

	cfg.Interactive = interactive
}

func setRsa(cfg *config.Config, field string, args []string) {
	switch field {
	case "public_key":
		cfg.Rsa.PublicKey = args[1]
	case "private_key":
		cfg.Rsa.PrivateKey = args[1]
	default:
		pterm.Warning.Printf("field %s is not found\n", args[0])
	}
}

func setHmac(cfg *config.Config, field string, args []string) {
	switch field {
	case "base64":
		base64, err := strconv.ParseBool(args[1])
		if err != nil {
			pterm.Fatal.Println("failed to parse base64 value")
		}

		cfg.Hmac.Base64Encoded = base64
	case "key":
		cfg.Hmac.Key = args[1]
	}

}

func setEcdsa(cfg *config.Config, field string, args []string) {
	switch field {
	case "public_key":
		cfg.Ecdsa.PublicKey = args[1]
	case "private_key":
		cfg.Ecdsa.PrivateKey = args[1]
	default:
		pterm.Warning.Printf("field %s is not found\n", args[0])
	}
}
