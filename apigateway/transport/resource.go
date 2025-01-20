package transport

import (
	"bytes"
	"io"
	"net/http"
	"xan-mortum/infrastructure/apigateway/config"
)

type ResourcesHttpClient struct {
	Config *config.Resources
}

func NewResourcesHttpClient(config *config.Resources) *ResourcesHttpClient {
	return &ResourcesHttpClient{
		Config: config,
	}
}

func (r *ResourcesHttpClient) GetBooks() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, r.Config.URL+"/api/v1/books", bytes.NewBuffer([]byte{}))
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

func (r *ResourcesHttpClient) GetUsers() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, r.Config.URL+"/api/v1/users", bytes.NewBuffer([]byte{}))

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
