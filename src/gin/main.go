package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type MyForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm MyForm
	c.Bind(&fakeForm)
	c.JSON(200, gin.H{
		"colors": fakeForm.Colors,
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

type Person struct {
	Name     string `form:"name"`
	Age      int    `form:"age"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func getForm(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name, person.Age, person.Address, person.Birthday)
	}

	c.JSON(200, gin.H{
		"name": person.Name,
		"age":  person.Age,
	})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// form
	r.GET("/form", indexHandler)
	r.POST("/form", formHandler)
	r.GET("/form/testing", getForm)

	// bind url
	r.GET("/user/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.JSON(200, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"name": person.Name,
			"age":  person.Age,
		})
	})

	// Sting
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// 静态文件
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 简单的路由组: v1
	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}

	// 简单的路由组: v2
	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Run("0.0.0.0:9099") // listen and serve on 0.0.0.0:8080
}
