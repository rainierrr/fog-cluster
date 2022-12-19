package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})

	// できるだけ環境変数からポート番号を取得する
	// 取得できない場合はデフォルトのポート番号を使用する

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	} else {
		log.Println("Server Run With Port ", port)
	}
}
