// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
)

// User models a Tick user.
type User struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	AvatarFileName string `json:"avatar_file_name"`
	Timezone       string `json:"timezone"`
	UpdatedAt      string `json:"updated_at"`
}

type Users []User

type UserStatus int

const (
	ActiveUsers UserStatus = iota
	DeletedUsers
)

func (c Client) GetUsers(ctx context.Context, status UserStatus) (Users, error) {
	users := Users{}
	var path string
	switch status {
	case ActiveUsers:
		path = "/users.json"
	case DeletedUsers:
		path = "/users/deleted.json"
	}
	err := c.get(ctx, path, &users)
	return users, err
}
