package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *ClientConfig) GetProjects() (map[string]int, error) {
	endpoint := fmt.Sprintf("%s/projects?token=%s", c.BaseURL, c.ApiToken)
	projects := new([]project)
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("request failed")
	}
	if err := json.NewDecoder(res.Body).Decode(projects); err != nil {
		return nil, fmt.Errorf("failed to decoded projects list")
	}
	processed := make(map[string]int)
	for _, p := range *projects {
		processed[p.Name] = p.ID
	}
	return processed, nil
}
