package main

import (
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

// Can be changed for user preference
var Duration time.Duration

func Monitor(CONN_STATUS *bool) {
	// Check the underlying ethernet interface

	for {
		checkInteface(CONN_STATUS)
		time.Sleep(Duration)
	}

}

func checkInteface(PREV_CONN_STATUS *bool) {

	grep := exec.Command("grep", "ethernet")
	list := exec.Command("nmcli", "device", "status")

	pipe, _ := list.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe

	list.Start()

	stdout, err := grep.Output()
	if err != nil {
		slog.Error("Error occured in Finding the Interface", "Error", err.Error())
	}

	res := string(stdout)

	if strings.Contains(res, "unavailable") { // Plugged out
		slog.Info("Ethernet Interface is  Unavailable")
		if *(PREV_CONN_STATUS) {
			*(PREV_CONN_STATUS) = false
		}
	} else if strings.Contains(res, "disconnected") { // similar to ethernet connection is set to off
		slog.Info("Ethernet Interface is Available but Disconnected")
		if *(PREV_CONN_STATUS) {
			*(PREV_CONN_STATUS) = false
		}
	} else {
		if checkPing() {
			slog.Info("Ethernet Inteface is Available,Connected and a Working Internet")
			if !*(PREV_CONN_STATUS) {
				*(PREV_CONN_STATUS) = true
				slog.Debug("Attempting to send Mail to notify")
				go SendMail()
			}
		} else {
			*(PREV_CONN_STATUS) = false
			slog.Info("Ethernet Interface is Available,Connected but a Working Internet")
		}

	}

}

// This function try to ping google.com
// some times there may be connection avaialble but no internet
func checkPing() bool {
	cmd := exec.Command("ping", "google.com", "-c", "2", "-w", "2")
	_, err := cmd.Output()
	return err == nil
}
