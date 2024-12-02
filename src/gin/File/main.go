package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SingleUploadFile(c *gin.Context) {

	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func MultiUploadFile(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件至指定目录
		dst := "./" + file.Filename
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func main() {
	router := gin.Default()

	router.POST("/uploadOne", func(c *gin.Context) {
		router.MaxMultipartMemory = 8 << 20 // 8 MiB
		SingleUploadFile(c)
	})

	router.POST("/uploadMulti", func(c *gin.Context) {
		router.MaxMultipartMemory = 8 << 20 // 8 MiB
		MultiUploadFile(c)
	})
	errorRun := router.Run("0.0.0.0:9099")
	if errorRun != nil {
		log.Println(errorRun)
	}
}
