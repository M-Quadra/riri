package test

import (
	"errors"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/M-Quadra/kazaana/v2"
	"github.com/M-Quadra/riri"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	pathBodyFormData0 = "/test/body/form-data/0"
	imgName           = "id85454797.jpg"
)

func init() {
	router.POST(pathBodyFormData0, func(c *gin.Context) {
		imgData, err := func(key string) ([]byte, error) {
			fileHeader, err := c.FormFile(key)
			if err != nil {
				return nil, err
			}

			f, err := fileHeader.Open()
			if err != nil {
				return nil, err
			}
			defer f.Close()

			return io.ReadAll(f)
		}("img")
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}

		err = func() error {
			imgDataOrg, err := os.ReadFile(imgName)
			if err != nil {
				return err
			}

			err = errors.New("different image")

			if len(imgData) != len(imgDataOrg) {
				return err
			}

			for i, v := range imgData {
				if v != imgDataOrg[i] {
					return err
				}
			}

			return nil
		}()
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}

		txt := c.PostForm("txt")
		if txt != "txt" {
			c.String(http.StatusOK, "Text error")
		}

		c.String(http.StatusOK, "1")
	})
}

func TestFormData(t *testing.T) {
	imgData, err := os.ReadFile(imgName)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	RunRouter()

	resData, kerr := localHost.Path(pathBodyFormData0).POST().
		Body.FormData(map[string]riri.FormDataValue{
		"img": {FileName: imgName, FileData: imgData},
		"txt": {Text: "txt"},
	}).Result()

	assert.False(t, kerr.HasError())
	assert.Equal(t, "1", string(resData))
}
