package config

type Config struct {
	JWT JWT
}

type JWT struct {
	Algorithms  []string
	Expirations []string
}

func Load() Config {
	return Config{
		JWT: JWT{
			Algorithms: []string{
				"HS256",
				"HS384",
				"HS512",
				"RS256",
				"RS384",
				"RS512",
				"ES256",
				"ES384",
				"ES512",
				"PS256",
				"PS384",
				"PS512",
			},
			Expirations: []string{
				"5m",
				"10m",
				"1h",
				"24h",
				"720h",
				"8760h",
			},
		},
	}
}
