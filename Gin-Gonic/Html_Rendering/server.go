package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate_func(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	// // 1. use LoadHTMLGlob or LoadHTMLFiles
	// // /index is not working...
	// router.LoadHTMLGlob("templates/**/*")
	// //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// // 2. use template.Must
	// html := template.Must(template.ParseFiles("templates/index.tmpl", "templates/index/index.tmpl", "templates/posts/index.tmpl", "templates/users/index.tmpl"))
	// router.SetHTMLTemplate(html)

	// // 3. use Delims
	// // /index is not working...
	// router.Delims("{{", "}}")
	// router.LoadHTMLGlob("templates/**/*")
	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })
	// router.GET("/index/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })
	// router.GET("/posts/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
	// 		"title": "Posts",
	// 	})
	// })
	// router.GET("/users/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	// 		"title": "Users",
	// 	})
	// })

	// 4. custom template
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate_func,
	})
	router.LoadHTMLFiles("./testdata/template/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2024, 05, 07, 0, 9, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
