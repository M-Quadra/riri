package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/M-Quadra/kazaana/v2"
	"github.com/M-Quadra/riri"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	pathTestBinary = "/testBinary"
)

func init() {
	router.POST(pathTestBinary, func(c *gin.Context) {
		data, err := c.GetRawData()
		if kazaana.HasError(err) {
			c.JSON(http.StatusOK, resInfo{
				Code: 0,
				Msg:  err.Error(),
			})
		}
		if string(data) != "ok" {
			c.JSON(http.StatusOK, resInfo{
				Code: 0,
				Msg:  "input error",
			})
			return
		}

		c.JSON(http.StatusOK, resInfo{
			Code: 1,
			Msg:  "sucess",
		})
	})
}

func TestBinary(t *testing.T) {
	RunRouter()

	info := resInfo{}
	resData, kerr := riri.POST(host + pathTestBinary).
		Body.Binary([]byte("ok")).
		BindJSON(&info)
	fmt.Println(string(resData))

	assert.False(t, kerr.HasError())
	assert.Equal(t, 1, info.Code)
}
