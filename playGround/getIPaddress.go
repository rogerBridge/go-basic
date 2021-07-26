package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := strings.Split(conn.LocalAddr().String(), ":")
	fmt.Println(localAddr[0])
}

func main() {
	GetOutboundIP()
}
