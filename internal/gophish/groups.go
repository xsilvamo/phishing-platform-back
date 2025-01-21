package gophish

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

// GetGroupByID obtiene un grupo específico por su ID desde GoPhish
func (s *GroupService) GetGroupByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups/%d", s.baseURL, id)

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
			if message, ok := errorResponse["message"].(string); ok && message == "Group not found" {
				return nil, fmt.Errorf("grupo no encontrado")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var group map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return group, nil
}

// GetGroupsSummary obtiene un resumen de todos los grupos desde GoPhish
func (s *GroupService) GetGroupsSummary() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups/summary", s.baseURL)

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

	var summaries map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&summaries); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return summaries, nil
}

// GetGroupSummaryByID obtiene un resumen de un grupo específico por su ID desde GoPhish
func (s *GroupService) GetGroupSummaryByID(id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups/%d/summary", s.baseURL, id)

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
			if message, ok := errorResponse["message"].(string); ok && message == "Group not found" {
				return nil, fmt.Errorf("grupo no encontrado")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var summary map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return summary, nil
}

// CreateGroup crea un nuevo grupo en GoPhish
func (s *GroupService) CreateGroup(data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups/", s.baseURL)

	// Codificar los datos del grupo
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

	var group map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return group, nil
}

// UpdateGroup modifica un grupo existente en GoPhish
func (s *GroupService) UpdateGroup(id int, data map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/groups/%d", s.baseURL, id)

	// Agregar el ID al payload
	data["id"] = id

	// Codificar los datos del grupo
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

		// Manejar errores específicos de GoPhish
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			if message, ok := errorResponse["message"].(string); ok && message == "Group not found" {
				return nil, fmt.Errorf("grupo no encontrado")
			}
		}

		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var updatedGroup map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&updatedGroup); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return updatedGroup, nil
}

// DeleteGroup elimina un grupo existente en GoPhish
func (s *GroupService) DeleteGroup(id int) error {
	url := fmt.Sprintf("%s/api/groups/%d", s.baseURL, id)

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
			if message, ok := errorResponse["message"].(string); ok && message == "Group not found" {
				return fmt.Errorf("grupo no encontrado")
			}
		}

		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	return nil
}

// ImportGroup procesa un archivo CSV y devuelve los datos como una lista de objetivos
func (s *GroupService) ImportGroup(filePath string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/api/import/group", s.baseURL)

	// Crear el archivo multipart para la carga
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("error creando el formulario del archivo: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("error copiando el archivo: %v", err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var targets []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&targets); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return targets, nil
}
