package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aakaru/gollback/config"
	"io"
	"net/http"
)

type Client struct {
	config     *config.Config
	httpClient *http.Client
}

type Workflow struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Active    bool                   `json:"active"`
	CreatedAt string                 `json:"createdAt"`
	UpdatedAt string                 `json:"updatedAt"`
	Nodes     []interface{}          `json:"nodes"`
	RawData   map[string]interface{} `json:"-"`
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		config:     cfg,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetWorkflows() ([]Workflow, error) {
	url := c.config.N8nURL + "/workflows"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("X-N8N-API-KEY", c.config.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v, err")
	}
	var response struct {
		Data []Workflow `json:"data"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	return response.Data, nil
}

func (c *Client) GetWorkflowByID(id string) (map[string]interface{}, error) {
	url := c.config.N8nURL + "/workflows/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("X-N8N-API-KEY", c.config.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v, err", err)
	}

	var workflow map[string]interface{}
	if err := json.Unmarshal(body, &workflow); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	return workflow, nil
}

func (c *Client) CreateWorkflow(workflow map[string]interface{}) error {
	url := c.config.N8nURL + "/workflows"
	jsonData, err := json.Marshal(workflow)
	if err != nil {
		return fmt.Errorf("failed to marshal workflow: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	req.Header.Set("X-N8N-API-KEY", c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
