package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "john@mail.com"

	user, err := NewUser(name, email, "password")
	assert.Nil(t, err)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.ID)
	assert.Nil(t, user.CheckPassword("password"))
	assert.NotNil(t, user.CheckPassword("wrong password"))
}
