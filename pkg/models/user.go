package models

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
}

func (u *User) Save() error {
	// Logic to save the user data to the database
	return nil
}

func (u *User) Update() error {
	// Logic to update the user data in the database
	return nil
}

func (u *User) Delete() error {
	// Logic to delete the user data from the database
	return nil
}

func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

var Users = []User{
	{
		ID:       1,
		Username: "john_doe",
		Email:    "jd@test.com",
	},
	{
		ID:       2,
		Username: "jane_doe",
		Email:    "jd2@test.com",
	},
}
