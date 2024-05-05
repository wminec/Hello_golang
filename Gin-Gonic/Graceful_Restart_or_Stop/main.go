package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 서비스 접속
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 5초의 타임아웃으로 인해 인터럽트 신호가 서버를 정상종료 할 때까지 기다립니다.
	quit := make(chan os.Signal)
	// kill (파라미터 없음) 기본값으로 syscall.SIGTERM 를 보냅니다.
	// kill -2 는 syscall.SIGINT 를 보냅니다.
	// kill -9 는 syscall.SIGKILL 를 보내지만 캐치할 수 없으므로, 추가할 필요가 없습니다.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 여기서 신호가 오는 것을 기다림
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 5초의 타임아웃으로 ctx.Done()을 캐치합니다.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 secounds.")
	}
	log.Println("Server exiting")
}
