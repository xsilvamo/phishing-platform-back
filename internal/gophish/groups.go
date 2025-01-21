package gophish

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GroupService maneja la lógica relacionada con grupos en GoPhish
type GroupService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewGroupService inicializa un nuevo GroupService
func NewGroupService(client *http.Client, apiKey, baseURL string) *GroupService {
	return &GroupService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetGroups obtiene todos los grupos desde GoPhish
func (s *GroupService) GetGroups() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups", s.baseURL)

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

	var groups []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return groups, nil
}
