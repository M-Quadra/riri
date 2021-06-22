package test

import (
	"testing"

	"github.com/M-Quadra/riri"
)

func TestRequest(t *testing.T) {
	_, kerr := riri.GET("123").Result()
	if !kerr.CheckError() {
		t.Fail()
	}
}
