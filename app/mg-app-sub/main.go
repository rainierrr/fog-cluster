package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	for {
		percent, err := cpu.Percent(time.Second*3, true)
		if err != nil {
			log.Println("error: ", err)
		}

		f, err := os.Create("./cpu.csv")
		if err != nil {
			log.Println("error: ", err)
		}

		var str_percent []string
		for _, p := range percent {
			str_percent = append(str_percent, strconv.FormatFloat(p, 'f', 2, 64))
		}

		w := csv.NewWriter(f)
		w.Write(str_percent)
		w.Flush()
	}
}
