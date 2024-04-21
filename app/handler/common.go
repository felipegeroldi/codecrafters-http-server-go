package handler

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/my_http"
)

func Index(c net.Conn) error {
	response := "HTTP/1.1 200 OK\r\n\r\n"

	_, err := c.Write([]byte(response))
	return err
}

func NotFound(c net.Conn) error {
	response := "HTTP/1.1 404 Not Found\r\n\r\n"
	_, err := c.Write([]byte(response))

	return err
}

func Echo(r *my_http.Request, c net.Conn) error {
	body := strings.Replace(r.Path, "/echo/", "", -1)

	response := my_http.Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusMsg:  "OK",
		Body:       body,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(body)),
		},
	}

	return response.WriteTo(c)
}

func UserAgent(r *my_http.Request, c net.Conn) error {
	var body string

	for h, v := range r.Headers {
		if h == "User-Agent" {
			body = v
		}
	}

	response := my_http.Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusMsg:  "OK",
		Body:       body,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(body)),
		},
	}

	return response.WriteTo(c)
}
