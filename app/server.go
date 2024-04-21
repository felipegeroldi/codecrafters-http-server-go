package main

import (
	"fmt"
	"net"
	"os"

	my_http "github.com/codecrafters-io/http-server-starter-go/app/http"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	HandleConnection(conn)
}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("Error occurred while reading data from connection, ", err.Error())
		os.Exit(1)
	}

	req := my_http.ParseData(buf)

	if req.Path == "/" {
		response := "HTTP/1.1 200 OK\r\n\r\n"
		if _, err := conn.Write([]byte(response)); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	} else {
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		if _, err := conn.Write([]byte(response)); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	}
}
