package main

import (
	"github.com/gin-gonic/gin"
	"log"
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

	r.Run("0.0.0.0:9090") // listen and serve on 0.0.0.0:8080
}
