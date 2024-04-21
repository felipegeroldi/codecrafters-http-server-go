package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/handler"
	"github.com/codecrafters-io/http-server-starter-go/app/my_http"
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
		if err := handler.Index(conn); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	} else if strings.HasPrefix(req.Path, "/echo/") {
		if err := handler.Echo(req, conn); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	} else if req.Path == "/user-agent" {
		if err := handler.UserAgent(req, conn); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	} else {
		if err := handler.NotFound(conn); err != nil {
			fmt.Println("Failed to write response, ", err.Error())
			os.Exit(1)
		}
	}
}
