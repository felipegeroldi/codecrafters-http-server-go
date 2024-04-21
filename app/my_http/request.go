package my_http

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
)

var (
	Bytes_WhiteSpace = []byte(" ")
	Bytes_CRLF       = []byte("\r\n")
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Headers  map[string]string
	Body     []byte
}

func ParseData(conn net.Conn) *Request {
	buf := make([]byte, 1024)
	headerData := make([]byte, 0)
	bodyData := make([]byte, 0)

	for {
		recvBytes, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error occurred while reading data from connection, ", err.Error())
			os.Exit(1)
		}

		if bytes.Contains(buf, []byte("\r\n\r\n")) {
			data := bytes.Split(buf[:recvBytes], []byte("\r\n\r\n"))
			headerData = append(headerData, data[0]...)
			bodyData = append(bodyData, data[1]...)

			break
		}

		headerData = append(headerData, buf...)
	}

	data := bytes.Split(headerData, Bytes_CRLF)

	reqData := bytes.Split(data[0], Bytes_WhiteSpace)
	data = data[1:]

	req := &Request{
		Method:   string(reqData[0]),
		Path:     string(reqData[1]),
		Protocol: string(reqData[2]),
		Headers:  make(map[string]string, 0),
	}

	var i int
	for i = 1; i < len(data); i++ {
		headerData := strings.Split(string(data[i]), ": ")
		req.Headers[string(headerData[0])] = string(headerData[1])
	}

	val, ok := req.Headers["Content-Length"]
	if ok {
		length, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		totalRecvBytes := len(bodyData)
		for {
			recvBytes, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error occurred while reading data from connection, ", err.Error())
				os.Exit(1)
			}

			bodyData = append(bodyData, buf[:recvBytes]...)

			totalRecvBytes += recvBytes
			if totalRecvBytes >= length {
				break
			}
		}

		req.Body = bytes.Trim(bodyData, "\x00")
	}

	return req
}
