package test

import (
	"math/rand"
	"testing"

	"github.com/M-Quadra/riri"
)

func TestGroup(t *testing.T) {
	tsGroup := riri.URLGroup{
		BaseURL: func() string {
			switch rand.Intn(2) {
			case 0:
				return "http://tsGroup/dev"
			case 1:
				return "http://tsGroup/debug"
			default:
				return "wtf"
			}
		},
	}

	switch tsGroup.Path("/ts") {
	case "http://tsGroup/dev/ts":
		return
	case "http://tsGroup/debug/ts":
		return
	default:
		t.Fail()
		break
	}
}
