package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ProfileService maneja la lógica relacionada con perfiles de envío
type ProfileService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewProfileService inicializa un nuevo ProfileService
func NewProfileService(client *http.Client, apiKey, baseURL string) *ProfileService {
	return &ProfileService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetProfiles obtiene todos los perfiles de envío desde GoPhish
func (s *ProfileService) GetProfiles() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/smtp", s.baseURL)

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

	var profiles []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profiles); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return profiles, nil
}

// GetProfileByID obtiene un perfil de envío específico por su ID
func (s *ProfileService) GetProfileByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/smtp/%d", s.baseURL, id)

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

	// Procesar respuesta en caso de error
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok {
				return nil, fmt.Errorf("error en la respuesta: %s", message)
			}
		}
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return profile, nil
}

// CreateProfile crea un nuevo perfil de envío en GoPhish
func (s *ProfileService) CreateProfile(data map[string]interface{}) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/api/smtp/", s.baseURL)

	// Convertir el payload a JSON
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error codificando datos: %v", err)
	}
	fmt.Println("Payload enviado:", string(payload))

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// Enviar la solicitud
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Manejar errores en la respuesta
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	// Decodificar la respuesta
	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return profile, nil
}

// UpdateProfile modifica un perfil de envío existente en GoPhish
func (s *ProfileService) UpdateProfile(id int, data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/smtp/%d", s.baseURL, id)

	// Agregar el campo "id" al payload
	data["id"] = id

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error codificando datos: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
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

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return profile, nil
}
