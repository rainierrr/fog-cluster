package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cache/persistence"
	// "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//store := persistence.NewInMemoryStore(10 * time.Second)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// router.GET("/cluster_metrics", cache.CachePage(store, 10*time.Second, getClusterUsageRateHandler))
	router.GET("/cluster_metrics", getClusterUsageRateHandler)
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
