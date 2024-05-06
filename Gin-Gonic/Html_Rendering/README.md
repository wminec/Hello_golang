# Html_Rendering

아래 내용을 학습할 수 있었음.
- html 템플릿 파일을 가져오는 방법
- go 템플릿을 이용하는 방법
- map 자료구조에 대한 학습

map 자료구조  
아래 코드에서 ```interface{}``` 는 모든 타입의 값을 가질 수 있는 타입. 따라서, 아래와 같이 사용할 수 있음.
```
data := map[string]interface{}{
    "name": "John Doe",
    "age":  30,
    "address": map[string]string{
        "city":    "San Francisco",
        "country": "USA",
    },
}
```

참고 : https://gin-gonic.com/ko-kr/docs/examples/html-rendering/