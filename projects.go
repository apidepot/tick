package gotick

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Project models a Tick project.
type Project struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Budget        float32 `json:"budget"`
	DateClosed    string  `json:"date_closed"`
	Notifications bool    `json:"notifications"`
	Billable      bool    `json:"billable"`
	Recurring     bool    `json:"recurring"`
	ClientID      uint    `json:"client_id"`
	OwnerID       uint    `json:"owner_id"`
	URL           string  `json:"url"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// FIXME(mdr): This should probably be a map[uint]Project using the Project.ID
// as the key
type Projects []Project

func GetOpenProjects(tickData JSONGetter) (Projects, error) {
	// FIXME(mdr) Get data in parallel instead of in series
	var allProjects Projects
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		projects, err := GetOpenProjectsOnPage(tickData, currentPage)
		if err != nil {
			return nil, err
		}
		if projects == Projects(nil) {
			foundLastPage = true
		} else {
			allProjects = append(allProjects, projects...)
			currentPage++
		}
	}
	return allProjects, nil
}

func GetOpenProjectsOnPage(tickData JSONGetter, page int) (Projects, error) {
	var projects Projects
	url := fmt.Sprintf("/projects.json?page=%d", page)
	data, err := tickData.GetJSON(url)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(data, []byte("")) {
		return nil, nil
	}
	err = json.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
