package test

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port = func() string {
		rand.Seed(time.Now().UTC().UnixNano())
		return ":9" + strconv.FormatInt(int64(rand.Intn(999)), 10)
	}()

	router = gin.Default()
	once   = sync.Once{}
)

const (
	url = "http://localhost"
)

func RunRouter() {
	once.Do(func() {
		go router.Run(port)
	})

}
