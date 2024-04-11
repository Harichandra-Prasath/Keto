package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendMail() error {
	Recievers := strings.Split(os.Getenv("TO_ADDR"), ",")

	msg := []byte(
		"Subject: Notification from Keto\r\n" +
			"\r\n" +
			"Ethernet Connection Established. Please Check your Computer\r\n")

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("PASSWORD"), "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("FROM_ADDR"), Recievers, msg)
	if err != nil {
		fmt.Println(err.Error())
	}

	return err

}
