package model_test

import (
	"testing"

	"github.com/AlexSH61/firstRestAPi/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptoPassword)
}
