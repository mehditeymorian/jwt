package model

import "strings"

type Encode struct {
	Algorithm     string
	Expiration    string
	PublicKeyPath string
	Issuer        string
	PayloadStr    string
	Payload       map[string]any
}

func (e *Encode) Execute() {
	payload := make(map[string]any)

	fields := strings.Split(e.PayloadStr, "\n")

	for _, field := range fields {
		keyValue := strings.Split(field, ":")
		payload[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
	}

	e.Payload = payload
}
