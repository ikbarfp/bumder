package handler_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/ikbarfp/bumder/internal/user/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHttp(t *testing.T) {
	assert.NotNil(t, handler.NewHttp(validator.New(), nil))
}

func TestRegister(t *testing.T) {

}

func TestProfile(t *testing.T) {

}
