// https://gin-gonic.com/ko-kr/docs/examples/bind-body-into-dirrerent-structs/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/a", SomeHandler1)
	r.POST("/b", SomeHandler2)
	r.Run(":8080")
}

func SomeHandler1(c *gin.Context) {
	//curl -X POST -H "Content-Type: application/json" -d '{"foo":"value"}' http://localhost:8080/a
	//// -> 정상적인 값이 나옴.
	//curl -X POST -H "Content-Type: application/json" -d '{"bar":"value"}' http://localhost:8080/a
	//// -> 정상적인 값이 안나옴.
	objA := formA{}
	objB := formB{}
	// 아래의 c.ShouldBind는 c.Request.Body를 소모하며, 재이용이 불가능합니다.
	if errA := c.ShouldBind(&objA); errA == nil {
		//fmt.Println(objA.Foo)
		c.String(http.StatusOK, `the body should be formA`)
		// c.Request.Body 가 EOF 이므로 에러가 발생합니다.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, "SomeHandler1 Hello")
	}
}

func SomeHandler2(c *gin.Context) {
	//curl -X POST -H "Content-Type: application/json" -d '{"foo":"value"}' http://localhost:8080/b
	//// -> 정상적인 값이 나옴.
	//curl -X POST -H "Content-Type: application/json" -d '{"bar":"value"}' http://localhost:8080/b
	//// -> 정상적인 값이 나옴.
	objA := formA{}
	objB := formB{}
	// c.Request.Body를 읽고 context에 결과를 저장합니다.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		//fmt.Println(objA.Foo)
		c.String(http.StatusOK, `the body should be formA`)
		// context 에 저장된 body를 읽어 재이용 합니다.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		//fmt.Println(objB.Bar)
		c.String(http.StatusOK, `the body should be formB`)
		// 다른 형식을 사용할 수도 있습니다.
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, "SomeHandler2 Hello!")
	}
}
