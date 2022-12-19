package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type postRequestBody struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Pressure    float64 `json:"pressure"`
	Location    string  `json:"location"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})

	router.POST("/post", func(ctx *gin.Context) {
		requestBody := postRequestBody{}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
			log.Println("Bad Request: ", err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": requestBody})
	})

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
