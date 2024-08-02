package web

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := []byte("123456")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, password)
	assert.NoError(t, err)

}
