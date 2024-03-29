// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
	"fmt"
	"log"
)

// Project models a Tick project.
type Project struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Budget        float64 `json:"budget"`
	DateClosed    string  `json:"date_closed"`
	Notifications bool    `json:"notifications"`
	Billable      bool    `json:"billable"`
	Recurring     bool    `json:"recurring"`
	ClientID      uint    `json:"client_id"`
	OwnerID       uint    `json:"owner_id"`
	URL           string  `json:"url"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	TotalHours    float64 `json:"total_hours,omitempty"`
}

type Projects []Project

type ProjectStatus int

const (
	OpenProjects ProjectStatus = iota
	ClosedProjects
)

func (c Client) GetProject(ctx context.Context, projectID int) (Project, error) {
	var project Project
	path := fmt.Sprintf("/projects/%d.json", projectID)
	err := c.get(ctx, path, &project)
	return project, err
}

func (c Client) GetProjects(ctx context.Context, status ProjectStatus) (Projects, error) {
	var allProjects Projects
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		projects, err := c.getProjectsOnPage(ctx, status, currentPage)
		if err != nil {
			log.Printf("error getting projects on page %d", currentPage)
			return nil, err
		}
		if len(projects) == 0 {
			foundLastPage = true
		} else {
			allProjects = append(allProjects, projects...)
			currentPage++
		}
	}
	return allProjects, nil
}

func (c Client) getProjectsOnPage(ctx context.Context, status ProjectStatus, page int) (Projects, error) {
	var path string
	switch status {
	case OpenProjects:
		path = fmt.Sprintf("/projects.json?page=%d", page)
	case ClosedProjects:
		path = fmt.Sprintf("/projects/closed.json?page=%d", page)
	}
	var projects Projects
	err := c.get(ctx, path, &projects)
	return projects, err
}
