package repository_test

import (
	"github.com/ikbarfp/bumder/internal/auth/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, repository.New())
}

func TestFindByMobilePhone(t *testing.T) {

}

func TestCreateAuth(t *testing.T) {

}

func TestGenerateToken(t *testing.T) {

}

func TestValidateToken(t *testing.T) {

}

func TestValidatePin(t *testing.T) {

}
