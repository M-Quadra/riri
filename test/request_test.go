package test

import (
	"testing"

	"github.com/M-Quadra/riri"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	_, kerr := riri.GET("123").Result()

	assert.True(t, kerr.CheckError())
}
