package main

import (
	"log/slog"
	"os"

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

	Initialse_Mail()

	// Assuming the user starts with connection, even if its not the case
	// it will be updated to false
	PREV_CONN_STATUS := true

	Monitor(&PREV_CONN_STATUS)
}
