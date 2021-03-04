package test

import (
	"net/http"
	"testing"

	"github.com/M-Quadra/riri"
	"github.com/gin-gonic/gin"
)

const (
	path0 = "/tsGet"
	path1 = "/tsGet1"
)

func init() {
	router.GET(path0, func(c *gin.Context) {
		if c.Query("1") == "2" {
			c.Writer.Write([]byte("1"))
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

	resData, kerr := riri.GET(url + port + path0 + "?1=1").
		Params.Set(map[string]string{
		"1": "2",
	}).Result()
	if kerr.HasError() {
		t.Fail()
		return
	}

	if string(resData) != "1" {
		t.Fail()
	}
}

func TestGet1(t *testing.T) {
	RunRouter()

	info := struct {
		Msg string `json:"msg"`
	}{}

	_, kerr := riri.GET(url + port + path1).BindJSON(&info)
	if kerr.HasError() || info.Msg != "wtf" {
		t.Fail()
		return
	}
}
