package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
<<<<<<< HEAD
	//store := persistence.NewInMemoryStore(3 * time.Second)
=======
	store := persistence.NewInMemoryStore(3 * time.Second)
>>>>>>> e6dbce53280d34da25d76e2895b39fb0e518991e

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

<<<<<<< HEAD
	// router.GET("/cluster_metrics", cache.CachePage(store, 3*time.Second, getClusterUsageRateHandler))
	router.GET("/cluster_metrics", getClusterUsageRateHandler)
=======
	router.GET("/cluster_metrics", cache.CachePage(store, 3*time.Second, getClusterUsageRateHandler))
>>>>>>> e6dbce53280d34da25d76e2895b39fb0e518991e
	router.GET("/cpu_usage_rate", getCPUUsageRateHandler)
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
