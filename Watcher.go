package main

import (
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

func Monitor() {
	// Check the underlying ethernet interface
	for {
		time.Sleep(1 * time.Second)
		checkInteface()
	}

}

func checkInteface() {
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

	if strings.Contains(res, "unavailable") {
		slog.Info("Ethernet Interface is  Unavailable")
	} else if strings.Contains(res, "unavailable") {
		slog.Info("Ethernet Interface is Available but Disconnected")
	} else {
		slog.Info("Ethernet Inteface is Available and Connected")
	}

}
