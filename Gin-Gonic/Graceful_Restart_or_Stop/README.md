# Graceful Restart or Stop
http 서버를 종료할 때, Graceful 하게 종료해야 할 필요가 있음. (서비스를 다 처리한 후 종료되어야 고객의 요청이 에러를 발생하지 않음)  

참고 |
- https://gin-gonic.com/ko-kr/docs/examples/graceful-restart-or-stop/
- https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go

## 사용된 기술  
- Goroutines
- Context

위 내용 테스트 한 것은 Test/Hello_Multi_Goroutines_with_Context 를 참고


