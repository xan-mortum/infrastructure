package transport

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"xan-mortum/infrastructure/apigateway/config"
	"xan-mortum/infrastructure/apigateway/internal/command"
)

type AuthHttpClient struct {
	Config *config.Auth
}

func NewAuthHttpClient(config *config.Auth) *AuthHttpClient {
	return &AuthHttpClient{
		Config: config,
	}
}

func (a *AuthHttpClient) SignIn() ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, a.Config.URL+"/api/v1/token", bytes.NewBuffer([]byte{}))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []byte{}, err
	}

	return body, nil
}

func (a *AuthHttpClient) CheckToken(token string) ([]byte, error) {
	cmd := command.CheckToken{
		Token: token,
	}
	reqBody, err := json.Marshal(cmd)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(http.MethodPost, a.Config.URL+"/api/v1/check_token", bytes.NewBuffer(reqBody))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
