package fogNode

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FogNodeController struct {
	db *gorm.DB
}

func NewFogNodeController(db *gorm.DB) *FogNodeController {
	return &FogNodeController{db: db}
}
func (r *FogNodeController) GetHandler(ctx *gin.Context) {
	fogNodeRepository := FogNodeRepository{}
	fogNodes, err := fogNodeRepository.FindAll(r.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "error": err.Error()})
		log.Println("Internal Server Error: ", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK", "fog_nodes": fogNodes})
}

func (r *FogNodeController) PostHandler(ctx *gin.Context) {
	var requestBody createParams
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
		log.Println("Bad Request: ", err.Error())
		return
	}

	fogNodeRepository := FogNodeRepository{}
	if err := fogNodeRepository.Create(r.db, &requestBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "error": err.Error()})
		log.Println("Internal Server Error: ", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
