package models

import "errors"

type User struct {
	Name     string `json:"name"`
	password password
	Email    string `json:"email"`
}

func NewUser(name, email, password string) (User, error) {
	validatedPassword, err := newPassword(password)

	if err != nil {
		return User{}, err
	}

	return User{Name: name, Email: email, password: validatedPassword}, nil
}

type password string

func newPassword(strPassword string) (password, error) {
	if len(strPassword) < 8 {
		return password(""), errors.New("Password too short")
	}

	return password(strPassword), nil
}
