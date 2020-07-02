package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type StoryCreateRequest struct {
	Name        string `json:"name"`
	ProjectID   int    `json:"project_id"`
	Deadline    string `json:"deadline"` // iso string
	Description string `json:"description"`
}

func (c *ClientConfig) StoryCreate(name, project, deadline, description string) error {
	endpoint := fmt.Sprintf("%s/stories?token=%s", c.BaseURL, c.ApiToken)
	req, err := json.Marshal(StoryCreateRequest{
		Name:        name,
		ProjectID:   projects[project],
		Deadline:    deadline,
		Description: description,
	})
	if err != nil {
		return fmt.Errorf("request encode fail")
	}
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return fmt.Errorf("request failed")
	}
	if res.StatusCode != 201 {
		return err
	}
	return nil
}
