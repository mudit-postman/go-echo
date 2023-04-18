package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/get", get)
	router.GET("/headers", headers)
	router.POST("/post", post)
	router.GET("/status/:code", status)
	router.GET("/response-headers", responseHeaders)

	router.Run("localhost:8080")
}
