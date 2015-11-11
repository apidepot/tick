package gotick

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Task models a Tick task associated with a particular project.
type Task struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Budget     float64 `json:"budget"`
	Position   uint    `json:"position"`
	ProjectID  uint    `json:"project_id"`
	DateClosed string  `json:"date_closed"`
	Billable   bool    `json:"billable"`
	URL        string  `json:"url"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	TotalHours float64 `json:"total_hours"`
	Project    Project
}

type Tasks []Task

func GetOpenTasks(tickData JSONGetter) (Tasks, error) {
	var allTasks Tasks
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		tasks, err := GetOpenTasksOnPage(tickData, currentPage)
		if err != nil {
			return nil, err
		}
		if tasks == Tasks(nil) {
			foundLastPage = true
		} else {
			allTasks = append(allTasks, tasks...)
			currentPage++
		}
	}
	return allTasks, nil
}

func GetOpenTasksOnPage(tickData JSONGetter, page int) (Tasks, error) {
	var tasks Tasks
	url := fmt.Sprintf("/tasks.json?page=%d", page)
	data, err := tickData.GetJSON(url)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(data, []byte("[]")) {
		return nil, nil
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTask uses Tick's API to return a specific Task identified by its ID.
func GetTask(tickData JSONGetter, id int) (Task, error) {
	var task Task
	url := fmt.Sprintf("/tasks/%d.json", id)
	data, err := tickData.GetJSON(url)
	if err != nil {
		return Task{}, err
	}
	if bytes.Equal(data, []byte("[]")) {
		return Task{}, nil
	}
	err = json.Unmarshal(data, &task)
	return task, nil
}
