package gophish

import (
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"os"
)

// GoPhishService es la estructura para manejar la conexión con la API de GoPhish
type GoPhishService struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// NewGoPhishService inicializa el servicio
func NewGoPhishService() *GoPhishService {
	return &GoPhishService{
		APIKey:  os.Getenv("GOPHISH_API_KEY"),
		BaseURL: os.Getenv("GOPHISH_API_URL"),
		Client:  &http.Client{},
	}
}

// ListCampaigns obtiene las campañas desde GoPhish
func (s *GoPhishService) ListCampaigns() ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/campaigns/", s.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: %s", string(body))
	}

	var campaigns []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&campaigns); err != nil {
		return nil, err
	}

	return campaigns, nil
}
