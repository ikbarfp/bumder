package handler_test

import (
	"github.com/ikbarfp/bumder/internal/feeds/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHttp(t *testing.T) {
	assert.NotNil(t, handler.NewHttp(nil))
}

func TestLike(t *testing.T) {

}

func TestPass(t *testing.T) {

}
