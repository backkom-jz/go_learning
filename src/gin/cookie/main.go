package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	if errorOne := router.SetTrustedProxies([]string{"127.0.0.1"}); errorOne != nil {
		log.Fatal(errorOne)
	}

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			fmt.Printf("cookie not set")
			c.SetCookie("gin_cookie", "test001", 3600, "/", "127.0.0.1:9099", false, true)
			cookie1, err1 := c.Cookie("gin_cookie")
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Println(cookie1)
			}
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	errorRun := router.Run("0.0.0.0:9099")
	if errorRun != nil {
		log.Println(errorRun)
	}
}
