// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
	"fmt"
)

// Task models a Tick task associated with a particular project.
type Task struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Budget     float64   `json:"budget"`
	Position   uint      `json:"position"`
	ProjectID  uint      `json:"project_id"`
	DateClosed string    `json:"date_closed"`
	Billable   bool      `json:"billable"`
	URL        string    `json:"url"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
	TotalHours float64   `json:"total_hours"`
	EntryInfo  EntryInfo `json:"entries"`
	Project    Project
}

type EntryInfo struct {
	Count     int    `json:"count"`
	URL       string `json:"url"`
	UpdatedAt string `json:"updated_at"`
}

type Tasks []Task

type TaskStatus int

const (
	OpenTasks TaskStatus = iota
	ClosedTasks
)

// GetTask uses Tick's API to return a specific Task identified by its ID.
func (c Client) GetTask(ctx context.Context, id int) (Task, error) {
	var task Task
	path := fmt.Sprintf("/tasks/%d.json", id)
	err := c.get(ctx, path, &task)
	return task, err
}

// GetTasks uses Tick's API to return all open or closed tasks.
func (c Client) GetTasks(ctx context.Context, status TaskStatus) (Tasks, error) {
	var allTasks Tasks
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		tasks, err := c.getTasksOnPage(ctx, status, currentPage)
		if err != nil {
			return nil, err
		}
		if len(tasks) == 0 {
			foundLastPage = true
		} else {
			allTasks = append(allTasks, tasks...)
			currentPage++
		}
	}
	return allTasks, nil
}

// GetProjectTasks uses Tick's API to return all open or closes tasks for a
// particular project.
func (c Client) GetProjectTasks(ctx context.Context, status TaskStatus, project string) (Tasks, error) {
	var tasks Tasks
	path := pathProjectsTasks(status, project)
	err := c.get(ctx, path, &tasks)
	return tasks, err
}

func (c Client) getTasksOnPage(ctx context.Context, status TaskStatus, page int) (Tasks, error) {
	var tasks Tasks
	path := pathTasksOnPage(status, page)
	err := c.get(ctx, path, &tasks)
	return tasks, err
}

func pathProjectsTasks(status TaskStatus, project string) string {
	if status == ClosedTasks {
		return fmt.Sprintf("/projects/%s/tasks/closed.json", project)
	}
	return fmt.Sprintf("/projects/%s/tasks.json", project)
}

func pathTasksOnPage(status TaskStatus, page int) string {
	if status == ClosedTasks {
		return fmt.Sprintf("/tasks/closed.json?page=%d", page)
	}
	return fmt.Sprintf("/tasks.json?page=%d", page)
}
