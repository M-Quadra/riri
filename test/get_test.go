package test

import (
	"net/http"
	"testing"

	"github.com/M-Quadra/riri"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	path0 = "/tsGet"
	path1 = "/tsGet1"
)

func init() {
	router.GET(path0, func(c *gin.Context) {
		if c.Query("1") == "2" {
			c.String(http.StatusOK, "1")
		}
	})

	router.GET(path1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "wtf",
		})
	})
}

func TestGet(t *testing.T) {
	RunRouter()

	resData, kerr := riri.GET(host + path0 + "?1=1").
		Params.Set(map[string]string{
		"1": "2",
	}).Result()

	assert.False(t, kerr.HasError())
	assert.Equal(t, "1", string(resData))
}

func TestGet1(t *testing.T) {
	RunRouter()

	info := struct {
		Msg string `json:"msg"`
	}{}

	_, kerr := riri.GET(host + path1).BindJSON(&info)

	assert.False(t, kerr.HasError())
	assert.Equal(t, "wtf", info.Msg)
}
