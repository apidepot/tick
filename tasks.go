// Copyright (c) 2016-2022 The tick developers. All rights reserved.
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

func (c Client) GetTasks(ctx context.Context, status TaskStatus) (Tasks, error) {
	var allTasks Tasks
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		tasks, err := c.getTasksOnPage(ctx, status, currentPage)
		if err != nil {
			return nil, err
		}
		if tasks == nil {
			foundLastPage = true
		} else {
			allTasks = append(allTasks, tasks...)
			currentPage++
		}
	}
	return allTasks, nil
}

func (c Client) getTasksOnPage(ctx context.Context, status TaskStatus, page int) (Tasks, error) {
	var path string
	switch status {
	case OpenTasks:
		path = fmt.Sprintf("/tasks.json?page=%d", page)
	case ClosedTasks:
		path = fmt.Sprintf("/tasks/closed.json?page=%d", page)
	}
	var tasks Tasks
	err := c.get(ctx, path, &tasks)
	return tasks, err
}
