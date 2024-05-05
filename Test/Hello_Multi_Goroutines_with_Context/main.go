package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("First goroutine cancelled")
		case <-time.After(10 * time.Second):
			fmt.Println("First goroutine finished")
		}
	}()

	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("Second goroutine cancelled")
		case <-time.After(10 * time.Second):
			fmt.Println("Second goroutine finished")
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT)
	fmt.Println("Press Ctrl+C to cancel first goroutines")
	select {
	// 여기서 신호가 오는 것을 기다림
	case sig := <-quit:
		switch sig {
		case syscall.SIGINT:
			cancel()
		case syscall.SIGQUIT:
			cancel2()
		}
	}

	// Wait for goroutines to finish
	fmt.Println("Waiting for goroutines to finish")
	time.Sleep(5 * time.Second)
}
