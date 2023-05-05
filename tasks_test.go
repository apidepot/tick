// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestGetOpenTasks(t *testing.T) {
	t.Skip()
	b, err := os.ReadFile(filepath.Join("testdata", "tasks_page1.json"))
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{
		Transport: fakeService(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(bytes.NewReader(b)),
			}, nil
		}),
	}
	tick, err := NewClient("mytoken", "subID", "thusUserAgent", WithHTTPClient(client))
	if err != nil {
		t.Fatalf("Error creating client for user testing: %s", err)
	}

	data, err := tick.GetTasks(context.TODO(), OpenTasks)
	if err != nil {
		t.Fatalf("Error getting tasks: %s", err)
	}
	if len(data) != 6 {
		t.Errorf("wrong len data for tasks / got : %d / want: TBD", len(data))
	}
}

func TestGetOneTask(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("testdata", "singletask.json"))
	if err != nil {
		t.Fatal(err)
	}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", b)
	}))
	defer svr.Close()

	tick, err := NewClient("mytoken", "subID", "thusUserAgent", WithBaseURL(svr.URL))
	if err != nil {
		t.Fatalf("Error creating client for user testing: %s", err)
	}

	taskID := 25
	task, err := tick.GetTask(context.TODO(), taskID)
	if err != nil {
		t.Fatalf("Error getting tasks: %s", err)
	}

	if int(task.ID) != taskID {
		t.Errorf("wrong task ID / got = %d / want %d", task.ID, taskID)
	}

	if task.Name != "Install exhaust port" {
		t.Errorf("wrong task name / got = %s", task.Name)
	}
}
