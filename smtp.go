package main

import (
	"log/slog"
	"net/smtp"
	"os"
	"strings"
)

var AUTH smtp.Auth
var MSG []byte
var RECIEVERS []string
var SMTP_USER string
var PASSWORD string
var FROM_ADDR string

func Initialse_Mail() {
	RECIEVERS = strings.Split(os.Getenv("TO_ADDR"), ",")
	SMTP_USER = os.Getenv("SMTP_USER")
	PASSWORD = os.Getenv("PASSWORD")
	FROM_ADDR = os.Getenv("FROM_ADDR")
	MSG = []byte(
		"Subject: Notification from Keto\r\n" +
			"\r\n" +
			"Ethernet Connection Established. Please Check your Computer\r\n")

	AUTH = smtp.PlainAuth("", SMTP_USER, PASSWORD, "smtp.gmail.com")
}

func SendMail() {
	// wrapper to execute as a seperate go routine
	// so that the delay wont be affected

	err := smtp.SendMail("smtp.gmail.com:587", AUTH, FROM_ADDR, RECIEVERS, MSG)
	if err != nil {
		slog.Error("Error sending Mail", "error", err.Error())
	} else {
		slog.Info("Successfully Notified the user via Mail")
	}
}
