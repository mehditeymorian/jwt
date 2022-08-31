package decode

import (
	"encoding/json"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mehditeymorian/jwt/internal/cmd"
	"github.com/mehditeymorian/jwt/internal/config"
	"github.com/mehditeymorian/jwt/internal/jwt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:     "decode",
		Short:   "decode jwt token",
		Long:    "decode jwt token",
		Example: "jwt decode <TOKEN>",
		Run:     decode,
	}

	return command
}

func decode(c *cobra.Command, args []string) {
	configPath := cmd.GetConfigPath(c)

	cfg := config.Load(configPath)

	strToken := ""

	if len(args) > 0 {
		strToken = args[0]
	} else {
		prompt := &survey.Input{Message: "JWT Token"}

		survey.AskOne(prompt, &strToken)

	}

	token, err := jwt.Decode(strToken, cfg)
	if err != nil {
		pterm.Fatal.Printf("failed to decode token: %v", err)
	}

	if token.Valid {
		pterm.Success.Println("Token is Valid.")
	} else {
		pterm.Warning.Println("Token is Invalid.")
	}

	result := make(map[string]any)
	result["headers"] = token.Header
	result["payload"] = token.Claims

	indent, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		pterm.Fatal.Printf("error marshaling result: %v", err)
	}

	indent = pretty.Color(indent, nil)

	pterm.DefaultBox.WithTitle("Decoded JWT Token").Println(string(indent))
}
