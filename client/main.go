package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	requestAndResponse(conn, "request0")
	requestAndResponse(conn, "request1")
	requestAndResponse(conn, "request2")

	requestAndResponse(conn, "close")
	time.Sleep(time.Hour)
}

func requestAndResponse(conn net.Conn, requestData string) {
	requestByteData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(requestByteData)

	responseData := make([]byte, 100)
	n, err := conn.Read(responseData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(responseData[:n]))
}
