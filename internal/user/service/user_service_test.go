package service_test

import (
	"github.com/ikbarfp/bumder/internal/user/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, service.New(nil, nil))
}

func TestRegister(t *testing.T) {

}

func TestGetProfile(t *testing.T) {

}
