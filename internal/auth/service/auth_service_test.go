package service_test

import (
	"github.com/ikbarfp/bumder/internal/auth/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, service.New(nil))
}

func TestLogin(t *testing.T) {

}
