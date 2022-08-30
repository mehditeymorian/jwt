package config

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/knadh/koanf"
	koanfYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v3"
)

const (
	PREFIX = "JWT"
	Name   = "jwt"
	Dir    = "/etc/jwt/"
	Path   = "/etc/jwt/config.yaml"
)

type Config struct {
	k           *koanf.Koanf
	loadedPath  string
	Algorithms  []string
	Expirations []string

	SigningMethod SigningMethod `koanf:"signing_method"`
	Algorithm     string        `koanf:"algorithm"`
	Rsa           *Rsa          `koanf:"rsa"`
	Hmac          *HMac         `koanf:"hmac"`
	Ecdsa         *Ecdsa        `koanf:"ecdsa"`
}

func Load(path string) *Config {
	var cfg Config

	// load configuration from file
	configPath := configFileAddress(path)
	log.Printf("reading config from %s\n", configPath)

	cfg.loadedPath = configPath

	k := koanf.New(".")
	cfg.k = k

	// load default configuration
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	if err := k.Load(file.Provider(configPath), koanfYaml.Parser()); err != nil {
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

	return &cfg
}

func (c *Config) Save() {
	log.Printf("saving config in %s\n", c.loadedPath)

	file, err := os.OpenFile(c.loadedPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating config file: %v", err)
	}
	defer file.Close()

	data := make(map[string]any)

	c.k.Unmarshal("", &data)

	encoder := yaml.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("failed to write config to file: %v\n", err)
	}
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}

func (c *Config) PrintableConfig() map[string]any {
	result := make(map[string]any)

	var config any
	switch c.SigningMethod {
	case RSA:
		config = c.Rsa
	case HMAC:
		config = c.Hmac
	case ECDSA:
		config = c.Ecdsa
	}

	result["signing_method"] = c.SigningMethod
	result["algorithm"] = c.Algorithm
	result[string(c.SigningMethod)] = config

	return result
}

func (c *Config) Print() {
	indent, err := json.MarshalIndent(c.PrintableConfig(), "", "\t")
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

}

func configFileAddress(userPath string) string {
	if len(userPath) != 0 {
		return userPath
	}

	dir, err := os.Getwd()
	if err != nil {
		return Path
	}

	pattern, err := regexp.Compile("jwt-config\\.ya*ml")
	if err != nil {
		return Path
	}

	dirFiles, err := os.ReadDir(dir)
	for _, entry := range dirFiles {
		if pattern.MatchString(entry.Name()) {
			return entry.Name()
		}
	}

	return Path
}
