package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CampaignService maneja la lógica relacionada con campañas en GoPhish
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

// GetCampaigns obtiene una lista de todas las campañas desde GoPhish
func (s *CampaignService) GetCampaigns() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns", s.baseURL)

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

// GetCampaignByID obtiene los detalles de una campaña específica por su ID
func (s *CampaignService) GetCampaignByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns/%d", s.baseURL, id)

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

		// Manejar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Campaign not found" {
				return nil, fmt.Errorf("campaña no encontrada")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var campaign map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return campaign, nil
}

// CreateCampaign crea una nueva campaña en GoPhish
func (s *CampaignService) CreateCampaign(data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns/", s.baseURL)

	// Codificar los datos de la campaña
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error codificando datos: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var campaign map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return campaign, nil
}

// GetCampaignResults obtiene los resultados de una campaña específica por su ID
func (s *CampaignService) GetCampaignResults(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns/%d/results", s.baseURL, id)

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

		// Manejar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Campaign not found" {
				return nil, fmt.Errorf("campaña no encontrada")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	// Decodificar como objeto JSON
	var results map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return results, nil
}

// GetCampaignSummary obtiene el resumen general de una campaña específica por su ID
func (s *CampaignService) GetCampaignSummary(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/campaigns/%d/summary", s.baseURL, id)

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

		// Manejar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Campaign not found" {
				return nil, fmt.Errorf("campaña no encontrada")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	// Decodificar como objeto JSON
	var summary map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return summary, nil
}

// DeleteCampaign elimina una campaña específica en GoPhish por su ID
func (s *CampaignService) DeleteCampaign(id int) error {
	url := fmt.Sprintf("%s/api/campaigns/%d", s.baseURL, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		// Manejar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Campaign not found" {
				return fmt.Errorf("campaña no encontrada")
			}
		}

		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	return nil
}

// CompleteCampaign marca una campaña específica como completada en GoPhish
func (s *CampaignService) CompleteCampaign(id int) error {
	url := fmt.Sprintf("%s/api/campaigns/%d/complete", s.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Manejo de errores en la respuesta
	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok {
				if message == "Error completing campaign" {
					return fmt.Errorf("la campaña no existe o no se puede completar")
				}
				return fmt.Errorf("error en la respuesta: %s", message)
			}
		}
		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	return nil
}
