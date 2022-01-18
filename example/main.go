package main

import (
	"log"
	"runtime"

	loadavg "github.com/msaf1980/go-loadavg"
)

func main() {
	la, err := loadavg.Parse()
	if err != nil {
		log.Fatal(err)
	}

	switch runtime.GOOS {
	case "linux":
		log.Printf("load averages: loadavg1 %.2f loadavg5 %.2f loadavg10 %.2f", la[0], la[1], la[2])
	default:
		log.Printf("load averages: loadavg1 %.2f loadavg5 %.2f loadavg15 %.2f", la[0], la[1], la[2])
	}
}
