# Hello Multi Goroutines with Context
SIGINT 를 받으면 첫 번째 goroutines 가 종료되도록 설정.

모든 goroutines 는 10초 후 작업이 완료 되었다는 메시지를 출력.

그 전에 goroutines 가 메시지를 출력하기 전, SIGINT 가 수신되면, 첫 번째 goroutines 를 cancel.

이후, 5초가 지나면 프로그램 종료

참고 : github copilot