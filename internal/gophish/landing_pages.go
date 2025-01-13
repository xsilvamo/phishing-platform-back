package gophish

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LandingPageService maneja la lógica relacionada con páginas de aterrizaje
type LandingPageService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewLandingPageService inicializa un nuevo LandingPageService
func NewLandingPageService(client *http.Client, apiKey, baseURL string) *LandingPageService {
	return &LandingPageService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetLandingPages obtiene todas las páginas de aterrizaje desde GoPhish
func (s *LandingPageService) GetLandingPages() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/pages", s.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var pages []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&pages); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return pages, nil
}

// GetLandingPageByID obtiene una página de aterrizaje específica por su ID desde GoPhish
func (s *LandingPageService) GetLandingPageByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/pages/%d", s.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		// Procesar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Page not found" {
				return nil, fmt.Errorf("página no encontrada")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var page map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return page, nil
}
