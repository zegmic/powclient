package client

import (
	"fmt"
	"net"
	"os"
)

func GetChallenge() string {
	srv := os.Getenv("SERVER")
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("challenge"))
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	return string(buf[:n])
}

func GetQuote(challenge string, nonce int) string {
	srv := os.Getenv("SERVER")
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	answer := fmt.Sprintf("solve %s %d", challenge, nonce)
	_, err = conn.Write([]byte(answer))
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	return string(buf[:n])
}
