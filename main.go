package main

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	// Create the log file
	f, err := os.OpenFile("/tmp/Keto.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Error In Creating Log", "Error", err.Error())
	}
	defer f.Close()

	// logging configuration
	logger := slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	min, _ := strconv.Atoi(os.Getenv("CHECK_PERIOD"))
	Duration = time.Minute * time.Duration(min)

	Initialse_Mail()

	// Assuming the user starts with out connection, even if its not the case
	// it will be updated to true
	PREV_CONN_STATUS := false

	Monitor(&PREV_CONN_STATUS)
}
