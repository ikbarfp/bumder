package handler_test

import (
	"github.com/ikbarfp/bumder/internal/auth/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHttp(t *testing.T) {
	assert.NotNil(t, handler.NewHttp(nil))
}

func TestLogin(t *testing.T) {

}
