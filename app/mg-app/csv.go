package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func readCpuCsv() ([]float64, error) {
	// ファイルを開く
	file, err := os.Open("cpu.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// csvを読み込む
	r := csv.NewReader(file)
	row, err := r.Read()
	if err != nil {
		return nil, err
	}

	var cpuList []float64
	// stringをfloatに変換
	for _, cpu := range row {
		float_cpu, err := strconv.ParseFloat(cpu, 64)
		if err != nil {
			return nil, err
		}
		cpuList = append(cpuList, float_cpu)
	}

	return cpuList, nil
}

func getCPUUsageRateHandler(ctx *gin.Context) {
	cpuList, err := readCpuCsv()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"cpu_usage_rate": cpuList, "message": "OK"})
}
