// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		expectedURL := "/projects/1775556.json"
		if r.URL.Path != expectedURL {
			t.Errorf("expected: %s / got: %s", expectedURL, r.URL.Path)
		}
		body := `{
  "id": 1775556,
  "name": "XX.YYY Test Project",
  "budget": null,
  "date_closed": null,
  "notifications": false,
  "billable": false,
  "recurring": false,
  "client_id": 293661,
  "owner_id": 221461,
  "url": "https://secure.tickspot.com/81326/api/v2/projects/1775556.json",
  "created_at": "2019-02-21T16:44:44.000-05:00",
  "updated_at": "2022-10-03T14:40:18.000-04:00",
  "total_hours": 4744.85,
  "tasks": {
    "count": 4,
    "url": "https://secure.tickspot.com/81326/api/v2/projects/1775556/tasks.json",
    "updated_at": "2022-10-03T14:40:18.000-04:00"
  },
  "client": {
    "id": 293661,
    "name": "MyClient",
    "archive": false,
    "url": "https://secure.tickspot.com/81326/api/v2/clients/293661.json",
    "updated_at": "2022-10-03T14:40:18.000-04:00"
  }
}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c, err := NewClient("fakeToken", "fakeID", "my-user-agent",
		WithBaseURL(server.URL),
	)
	if err != nil {
		t.Errorf("expected new client err to be nil, got: %v", err)
	}
	_, err = c.GetProject(context.Background(), 1775556)
	if err != nil {
		t.Errorf("expected response getting project 1775556 to be nil; got: %v", err)
	}
}

func TestGetOpenProjectsOnFirstPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		var body string
		switch r.URL.Path {
		case "/projects.json?page=1":
			body = `[
  {
    "id": 2237757,
    "name": "MY.002 Test Project Two",
    "budget": null,
    "date_closed": null,
    "notifications": false,
    "billable": true,
    "recurring": false,
    "client_id": 373699,
    "owner_id": 256228,
    "url": "https://secure.tickspot.com/81326/api/v2/projects/2237757.json",
    "created_at": "2022-06-13T16:52:12.000-04:00",
    "updated_at": "2022-08-24T14:45:18.000-04:00"
  },
  {
    "id": 2093433,
    "name": "MY.003 Test Project Three",
    "budget": null,
    "date_closed": null,
    "notifications": false,
    "billable": true,
    "recurring": false,
    "client_id": 278516,
    "owner_id": 221461,
    "url": "https://secure.tickspot.com/81326/api/v2/projects/2093433.json",
    "created_at": "2021-04-25T07:05:11.000-04:00",
    "updated_at": "2022-02-23T19:50:31.000-05:00"
  }
]`
		case "/projects.json?page=2":
			log.Println("Trying to get page 2 of projects")
			body = `[]`
		default:
			t.Errorf("expected to request /projects.json?page=1 or 2, got: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))
	defer server.Close()

	c, err := NewClient("fakeToken", "fakeID", "my-user-agent",
		WithBaseURL(server.URL),
	)
	if err != nil {
		t.Errorf("expected new client err to be nil got %v", err)
	}
	res, err := c.GetProjects(context.Background(), OpenProjects)
	if err != nil {
		t.Errorf("foo expected response err to be nil got %v", err)
	}
	if len(res) != 2 {
		t.Errorf("bar expected to receive two projects but only got %d", len(res))
	}
	if res[0].ID != 2237757 {
		t.Errorf("expected the first project to have a project ID of 2237757 but got %d", res[0].ID)
	}
}

func serverMock() *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc("/projects.json", handleOpenProjects)
	router.HandleFunc("/projects/closed.json", handleClosedProjects)

	srv := httptest.NewServer(router)
	return srv
}

func handleOpenProjects(w http.ResponseWriter, r *http.Request) {
	projectsPage1 := `[
	{
		"id":1111111,
		"name":"P.001 First Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111111.json",
		"created_at":"2014-08-06T09:39:44.000-04:00",
		"updated_at":"2014-08-26T18:19:31.000-04:00"
	},
	{
		"id":1111112,
		"name":"P.002 Second Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111112.json",
		"created_at":"2014-08-07T17:43:39.000-04:00",
		"updated_at":"2014-08-13T12:22:35.000-04:00"
	}
]`
	_, _ = w.Write([]byte(projectsPage1))
}

func handleClosedProjects(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("mock closed projects response"))
}
