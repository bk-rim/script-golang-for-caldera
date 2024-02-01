package util

import (
	"log"
	"os/exec"
)

func StopProcess(exeName string) {
	cmd := exec.Command("taskkill", "/F", "/IM", exeName)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error stopping process", err)
	}
}
