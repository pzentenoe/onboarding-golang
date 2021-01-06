package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       string   `json:"id"`
	UserName string   `json:"user_name"`
	Password string   `json:"password"`
	Persona  *Persona `json:"persona"`
}

func (u *User) SetUserName(username string) {
	u.UserName = username
}

func (u *User) ToJSON() (string, error) {
	dataBytes, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	return string(dataBytes), err
}

func (u *User) FromJSON(jsonData string) (*User, error) {
	var user *User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}

