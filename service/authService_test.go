package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAdmin(t *testing.T) {
	token, err := CreateToken("ADMIN", "", []byte(os.Getenv("SECRETKEY")))
	assert.Nil(t, err, "error in token creation: %v", err)
	err = CheckAdmin(token)
	assert.Nil(t, err, "error although the token is valid")

	token, err = CreateToken("USER", "", []byte(os.Getenv("SECRETKEY")))
	assert.Nil(t, err, "error in token creation")
	err = CheckAdmin(token)
	assert.NotNil(t, err, "The token is invalid, should cause an error")

}
