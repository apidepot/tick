package gotick

import "encoding/json"

// User models a Tick user.
type User struct {
	ID             uint64 `json:"id"`
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
