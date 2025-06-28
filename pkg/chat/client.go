package chat

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AiClient interface {
	GetResponse(cfg interface{}) (string, error)
}

type OpenaiClient struct {
	AiClientBasicComponents
}

func (c *OpenaiClient) GetResponse(cfg OpenaiApiConfig) (string, error) {
	jsonBody, err := json.Marshal(map[string]string{
		"model": cfg.Model,
		"input": cfg.Input,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, c.ApiEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

type DeepseekClient struct {
	AiClientBasicComponents
}

func (c *DeepseekClientStandardImpl)