package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/handler"
	"github.com/codecrafters-io/http-server-starter-go/app/my_http"
)

var directory *string

func init() {
	directory = flag.String("directory", "./", "Directory path for static files")
	flag.Parse()
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	req := my_http.ParseData(conn)

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
	} else if strings.HasPrefix(req.Path, "/files/") {
		if req.Method == my_http.METHOD_GET {
			if err := handler.GetFile(req, conn, *directory); err != nil {
				fmt.Println("Failed to write response, ", err.Error())
				os.Exit(1)
			}
		} else if req.Method == my_http.METHOD_POST {
			if err := handler.PostFile(req, conn, *directory); err != nil {
				fmt.Println("Failed to write response, ", err.Error())
				os.Exit(1)
			}
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
