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
	"github.com/pterm/pterm"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v3"
)

const (
	PREFIX = "JWT"
	Name   = "jwt"
	Dir    = "/etc/jwt/"
	Path   = "/etc/jwt/.jwt.yaml"
)

type Config struct {
	k           *koanf.Koanf
	loadedPath  string
	Algorithms  []string
	Expirations []string

	Interactive bool `koanf:"interactive"`

	Rsa   *Rsa   `koanf:"rsa"`
	Hmac  *HMac  `koanf:"hmac"`
	Ecdsa *Ecdsa `koanf:"ecdsa"`
}

type saveConfig struct {
	Interactive bool `koanf:"interactive"`

	Rsa   *Rsa   `koanf:"rsa" yaml:"rsa"`
	Hmac  *HMac  `koanf:"hmac" yaml:"hmac"`
	Ecdsa *Ecdsa `koanf:"ecdsa" yaml:"ecdsa"`
}

func Load(path string) *Config {
	var cfg Config

	// load configuration from file
	configPath := FileAddress(path)
	pterm.Info.Printf("reading config from %s\n", configPath)

	cfg.loadedPath = configPath

	k := koanf.New(".")
	cfg.k = k

	// load default configuration
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		pterm.Fatal.Printf("error loading default config: %v\n", err)
	}

	if err := k.Load(file.Provider(configPath), koanfYaml.Parser()); err != nil {
		pterm.Warning.Printf("error loading config.yaml: %v\n", err)
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
		pterm.Warning.Printf("error loading environment variables: %v\n", err)
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		pterm.Fatal.Printf("error unmarshaling config: %v\n", err)
	}

	return &cfg
}

func (c *Config) Save() {
	pterm.Info.Printf("saving config in %s\n", c.loadedPath)

	file, err := os.OpenFile(c.loadedPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		pterm.Fatal.Printf("error opening/creating config file: %v\n", err)
	}
	defer file.Close()

	saveCfg := saveConfig{
		Interactive: c.Interactive,
		Rsa:         c.Rsa,
		Hmac:        c.Hmac,
		Ecdsa:       c.Ecdsa,
	}

	encoder := yaml.NewEncoder(file)

	err = encoder.Encode(saveCfg)
	if err != nil {
		pterm.Fatal.Printf("failed to write config to file: %v\n", err)
	}
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}

func (c *Config) PrintableConfig() map[string]any {
	result := make(map[string]any)

	result["interactive"] = c.Interactive
	result[string(RSA)] = c.Rsa
	result[string(HMAC)] = c.Hmac
	result[string(ECDSA)] = c.Ecdsa

	return result
}

func (c *Config) Print() {
	indent, err := json.MarshalIndent(c.PrintableConfig(), "", "\t")
	if err != nil {
		pterm.Fatal.Printf("error marshal config: %v\n", err)
	}

	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(cfgStrTemplate, string(indent))

}

func (c *Config) PrintMode() {
	mode := "interactive mode"

	if !c.Interactive {
		mode = "option mode"
	}

	pterm.Info.Println(mode)
}

func FileAddress(userPath string) string {
	if len(userPath) != 0 {
		return userPath
	}

	dir, err := os.Getwd()
	if err != nil {
		return Path
	}

	pattern, err := regexp.Compile("\\.jwt\\.ya*ml")
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
