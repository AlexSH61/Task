package db_test

import (
	"testing"

	"github.com/AlexSH61/firstRestAPi/internal/app/db"
	"github.com/AlexSH61/firstRestAPi/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := db.TestDB(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
func TestUserRepo_FindByEmail(t *testing.T) {
	s, teardown := db.TestDB(t, databaseURL)
	defer teardown("users")

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.com",
	})
	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
