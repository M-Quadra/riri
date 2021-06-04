package test

import (
	"math/rand"
	"testing"

	"github.com/M-Quadra/riri"
	"github.com/gin-gonic/gin"
)

const (
	pathGroup0 = "/test/group/0"
)

func init() {
	router.GET(pathGroup0, func(c *gin.Context) {
	})
}

var (
	localHost = riri.Group{
		BaseURL: func() string {
			return "http://localhost" + port
		},
	}
)

func TestGroupPath(t *testing.T) {
	tsGroup := riri.Group{
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
	case "http://tsGroup/debug/ts":
	default:
		t.Fail()
	}
}

func TestGroupRequest(t *testing.T) {
	RunRouter()

	res, kerr := localHost.Path(pathGroup0).GET().Do()
	if kerr.HasError() {
		t.Fail()
		return
	}
	defer res.Body.Close()
}
