package riri

import (
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	path0 = "/tsGet"
)

func init() {
	router.GET(path0, func(c *gin.Context) {
		if c.Query("1") == "2" {
			c.Writer.Write([]byte("1"))
		}
	})
}

func TestGet(t *testing.T) {
	RunRouter()

	result, kerr := GET(url + port + path0 + "?1=1").Params.Set(map[string]string{
		"1": "2",
	}).Result()
	if kerr.HasError() {
		t.Fail()
		return
	}

	if string(result) != "1" {
		t.Fail()
	}
}
