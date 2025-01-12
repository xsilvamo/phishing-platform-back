package gophish

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CampaignService maneja la lógica relacionada con campañas de GoPhish
type CampaignService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewCampaignService inicializa un nuevo CampaignService
func NewCampaignService(client *http.Client, apiKey, baseURL string) *CampaignService {
	return &CampaignService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// ListCampaigns obtiene las campañas desde GoPhish
func (s *CampaignService) ListCampaigns() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns/", s.baseURL)

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

	var campaigns []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&campaigns); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return campaigns, nil
}
