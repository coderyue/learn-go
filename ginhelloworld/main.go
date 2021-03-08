package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/test", func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println("全路径:", fullPath)
		context.Writer.WriteString("success")
	})

	r.Run()
}
