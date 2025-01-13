package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TemplateService maneja la lógica relacionada con templates
type TemplateService struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

// NewTemplateService inicializa un nuevo TemplateService
func NewTemplateService(client *http.Client, apiKey, baseURL string) *TemplateService {
	return &TemplateService{
		client:  client,
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// GetTemplates obtiene todos los templates desde GoPhish
func (s *TemplateService) GetTemplates() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/templates", s.baseURL)

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

	var templates []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&templates); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return templates, nil
}

// GetTemplateByID obtiene un template específico por su ID desde GoPhish
func (s *TemplateService) GetTemplateByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/templates/%d", s.baseURL, id)

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
			if message, ok := errorResponse["message"].(string); ok && message == "Template not found" {
				return nil, fmt.Errorf("template no encontrado")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var template map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return template, nil
}

// CreateTemplate crea un nuevo template en GoPhish
func (s *TemplateService) CreateTemplate(data map[string]interface{}) (map[string]interface{}, error) {
	//print de validacion
	fmt.Println("Ingresando a la funcion CreateTemplate")
	url := fmt.Sprintf("%s/api/templates/", s.baseURL)

	payload, err := json.Marshal(data)
	fmt.Println("Data: ", data)
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

	var template map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return template, nil
}

// UpdateTemplate modifica un template existente en GoPhish
func (s *TemplateService) UpdateTemplate(id int, data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/templates/%d", s.baseURL, id)

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

		// Procesar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Template not found" {
				return nil, fmt.Errorf("template no encontrado")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var template map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return template, nil
}

// DeleteTemplate elimina un template existente en GoPhish
func (s *TemplateService) DeleteTemplate(id int) error {
	url := fmt.Sprintf("%s/api/templates/%d", s.baseURL, id)

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
			if message, ok := errorResponse["message"].(string); ok && message == "Template not found" {
				return fmt.Errorf("template no encontrado")
			}
		}

		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	return nil
}

// ImportEmail analiza un correo electrónico y devuelve su contenido como texto, HTML y asunto
func (s *TemplateService) ImportEmail(content string, convertLinks bool) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/import/email", s.baseURL)

	// Crear el payload
	payload := map[string]interface{}{
		"content":       content,
		"convert_links": convertLinks,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error codificando datos: %v", err)
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	// Decodificar la respuesta
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return result, nil
}
