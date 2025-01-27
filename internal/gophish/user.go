package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// UserService maneja la lógica relacionada con usuarios
type UserService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewUserService inicializa un nuevo servicio de usuarios
func NewUserService(client *http.Client, apiKey, baseURL string) *UserService {
	return &UserService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetCurrentUser obtiene la información del usuario autenticado desde GoPhish
func (s *UserService) GetCurrentUser() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/users/", s.baseURL)

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

	var users []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return users, nil
}

// CreateUser crea un nuevo usuario en GoPhish
func (s *UserService) CreateUser(data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/users/", s.baseURL)

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

	body, _ := io.ReadAll(resp.Body)

	// Manejar errores de la API
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok {
				return nil, fmt.Errorf("error de la API: %s", message)
			}
		}
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	// Procesar la respuesta si es válida
	var user map[string]interface{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return user, nil
}

// UpdateUser actualiza un usuario existente en GoPhish
func (s *UserService) UpdateUser(id int, data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/users/%d", s.baseURL, id)

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

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok {
				return nil, fmt.Errorf("error de la API: %s", message)
			}
		}
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var user map[string]interface{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return user, nil
}

// GetUsers obtiene todos los usuarios registrados en GoPhish
func (s *UserService) GetUsers() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/users/", s.baseURL)

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

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var users []map[string]interface{}
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return users, nil
}

// GetUserByID obtiene un usuario específico en GoPhish por su ID
func (s *UserService) GetUserByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/users/%d", s.baseURL, id)

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

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok {
				return nil, fmt.Errorf("error de la API: %s", message)
			}
		}
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var user map[string]interface{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return user, nil
}
