package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
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
	router.GET("/cpu_usage_rate", func(ctx *gin.Context) {
		// CPU使用率を取得
		percent, err := cpu.Percent(time.Millisecond*200, false)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"cpu_usage_rate": percent})
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
