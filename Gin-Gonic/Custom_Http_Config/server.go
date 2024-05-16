package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//func main() {
//	router := gin.Default()
//	http.ListenAndServe(":8080", router)
//}

func main() {
	router := gin.Default()

	s := &http.Server{
		// IP 주소와 포트 번호
		Addr: ":8080",
		// HTTP 요청을 처리하는 핸들러
		Handler: router,
		// 서버가 클라이언트의 요청을 읽는 데 걸리는 최대 시간
		ReadTimeout: 10 * time.Second,
		// 서버가 응답을 작성하는 데 걸리는 최대 시간
		WriteTimeout: 10 * time.Second,
		// 클라이언트의 HTTP 요청 헤더의 최대 크기
		MaxHeaderBytes: 1 << 20,
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	s.ListenAndServe()

}
