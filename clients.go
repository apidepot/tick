// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
	"fmt"
)

// TickClient models a client in Tick. Named TickClient to differentiate from
// Client, which is the REST API Client.
type TickClient struct {
	ID           uint         `json:"id"`
	Name         string       `json:"name"`
	Archive      bool         `json:"archive"`
	URL          string       `json:"url"`
	UpdatedAt    string       `json:"updated_at"`
	ProjectsInfo ProjectsInfo `json:"projects"`
}

type TickClients []TickClient

type ProjectsInfo struct {
	Count     int    `json:"count"`
	URL       string `json:"url"`
	UpdatedAt string `json:"updated_at"`
}

// GetClients uses Tick's API to return all clients with open projects.
func (c Client) GetClients(ctx context.Context) (TickClients, error) {
	var tickClients TickClients
	path := "/clients.json"
	err := c.get(ctx, path, &tickClients)
	return tickClients, err
}

// GetClient uses Tick's API to return the specified client with a summary of
// project information.
func (c Client) GetClient(ctx context.Context, client int) (TickClient, error) {
	var tickClient TickClient
	path := fmt.Sprintf("/clients/%d.json", client)
	err := c.get(ctx, path, &tickClient)
	return tickClient, err
}
