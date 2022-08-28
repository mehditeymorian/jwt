package config

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

const (
	PREFIX = "JWT"
	Name   = "jwt"
	Dir    = "/etc/jwt/"
	Path   = "/etc/jwt/config.yaml"
)

type Config struct {
	Algorithms  []string
	Expirations []string

	SigningMethod SigningMethod `koanf:"signing_method"`
	Algorithm     string        `koanf:"algorithm"`
	Rsa           Rsa           `koanf:"rsa"`
	Hmac          HMac          `koanf:"hmac"`
}

func Load() Config {
	var cfg Config

	k := koanf.New(".")

	// load default configuration
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider(Path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml: %v", err)
	}

	// load environment variables
	cb := func(key string, value string) (string, interface{}) {
		finalKey := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(key, PREFIX)), "__", ".")

		if strings.Contains(value, ",") {
			// remove all the whitespace from value
			// split the value using comma
			finalValue := strings.Split(removeWhitespace(value), ",")

			return finalKey, finalValue
		}

		return finalKey, value
	}
	if err := k.Load(env.ProviderWithValue(PREFIX, ".", cb), nil); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	indent, err := json.MarshalIndent(cfg.PrintableConfig(), "", "\t")
	if err != nil {
		log.Fatalf("error marshal config: %v", err)
	}

	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(cfgStrTemplate, string(indent))

	return cfg
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}

func (c Config) PrintableConfig() map[string]any {
	result := make(map[string]any)

	var config any
	switch c.SigningMethod {
	case RSA:
		config = c.Rsa
	case HMAC:
		config = c.Hmac
	}

	result["signing_method"] = c.SigningMethod
	result[string(c.SigningMethod)] = config

	return result
}
