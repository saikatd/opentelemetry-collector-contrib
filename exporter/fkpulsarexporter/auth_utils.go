package fkpulsarexporter

import (
	"encoding/base64"
	"encoding/json"
)

func GeneratePrivateKey(clientId, clientSecret, issuerUrl string) (string, error) {
	props := map[string]string{
		"type":          "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
		"issuer_url":    issuerUrl,
	}
	jsonStr, err := json.Marshal(props)
	if err != nil {
		return "", err
	}
	encodedText := base64.StdEncoding.EncodeToString(jsonStr)
	data := "data:application/json;base64," + encodedText
	return data, nil
}
