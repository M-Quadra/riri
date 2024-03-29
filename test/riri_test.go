package test

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/M-Quadra/kazaana/v2"
	"github.com/gin-gonic/gin"
)

var (
	port = func() string {
		rand.Seed(time.Now().UTC().UnixNano())
		return ":9" + strconv.FormatInt(int64(rand.Intn(999)), 10)
	}()

	host = "http://localhost" + port

	router = gin.Default()
	once   = sync.Once{}
)

func RunRouter() {
	go once.Do(func() {
		err := router.Run(port)
		if kazaana.HasError(err) {
			return
		}

		time.Sleep(5 * time.Millisecond)
	})
}

type resInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type resStr struct {
	resInfo
	Data string `json:"data,omitempty"`
}
