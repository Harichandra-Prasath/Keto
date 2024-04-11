package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	// Create the log file
	f, err := os.OpenFile("/tmp/Keto.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Error In Creating Log", "Error", err.Error())
	}
	defer f.Close()
	log.SetOutput(f)

	Monitor()
}
