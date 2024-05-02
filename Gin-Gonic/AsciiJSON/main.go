// https://gin-gonic.com/ko-kr/docs/examples/ascii-json/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO언어",
			"tag":  "<br>",
		}

		// 출력내용 : {"lang":"GO\uc5b8\uc5b4","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
