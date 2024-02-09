package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@emample.com",
		Password: "1234",
	}
}
