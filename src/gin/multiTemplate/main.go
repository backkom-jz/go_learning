package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "../templates/base.html", "../templates/index.html")
	r.AddFromFiles("article", "../templates/base.html", "../templates/index.html", "../templates/article.html")
	return r
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors": gin.H{
					"modelA": errA.Error(),
					"modelB": errB.Error(),
				},
			},
		)
	}
}

func main() {
	r := gin.Default()
	if errorOne := r.SetTrustedProxies([]string{"127.0.0.1"}); errorOne != nil {
		log.Fatal(errorOne)
	}

	r.HTMLRender = createMyRender()
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	r.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})
	r.POST("/bind", func(c *gin.Context) {
		SomeHandler(c)
	})

	if err := r.Run("127.0.0.1:9099"); err != nil {
		log.Fatal(err)
	}

}
