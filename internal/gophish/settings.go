package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SettingsService maneja la lógica relacionada con configuración de GoPhish
type SettingsService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewSettingsService inicializa un nuevo SettingsService
func NewSettingsService(client *http.Client, apiKey, baseURL string) *SettingsService {
	return &SettingsService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// ResetAPIKey resetea la clave API de GoPhish
func (s *SettingsService) ResetAPIKey() (string, error) {
	url := fmt.Sprintf("%s/api/reset", s.baseURL)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return "", fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %d", resp.StatusCode)
	}

	var response struct {
		Success bool   `json:"success"`
		Data    string `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return response.Data, nil
}
