package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexSH61/firstRestAPi/internal/app/model"
)

type APIClient struct {
	BaseURL string
}

func (c *APIClient) GetAllTasks() ([]model.Task, error) {
	resp, err := http.Get(fmt.Sprintf("%s/tasks", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []model.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *APIClient) CreateTask(newTask model.Task) (*model.Task, error) {
	jsonData, err := json.Marshal(newTask)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/tasks", c.BaseURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var createdTask model.Task
	if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
		return nil, err
	}

	return &createdTask, nil
}

func (c *APIClient) UpdateTaskStatus(updateData model.Task) error {
	jsonData, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks", c.BaseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("incorrect status task: %d", resp.StatusCode)
	}

	return nil
}

func (c *APIClient) DeleteTask(taskID int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/tasks/%d", c.BaseURL, taskID), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("incorrect status task: %d", resp.StatusCode)
	}

	return nil
}
