// Copyright (c) 2016 The gotick developers. All rights reserved.
// Project site: https://github.com/questrail/gotick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package gotick

import "encoding/json"

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

func GetUsers(tickData JSONGetter) (Users, error) {
	var users Users
	data, err := tickData.GetJSON("/users.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
