package main

import (
	"log"
	"mg-app/db"
	fogNode "mg-app/fog_node"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.AutoMigration()
	db := db.GetDB()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	fogNodeController := fogNode.NewFogNodeController(db)
	router.GET("/fog_nodes", fogNodeController.GetHandler)
	router.POST("/fog_node", fogNodeController.PostHandler)

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
