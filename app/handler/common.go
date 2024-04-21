package handler

import (
	"fmt"
	"log"
	"net"
	"os"
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

func Created(c net.Conn) error {
	response := "HTTP/1.1 201 Created\r\n\r\n"
	_, err := c.Write([]byte(response))

	return err
}

func Echo(r *my_http.Request, c net.Conn) error {
	body := strings.Replace(r.Path, "/echo/", "", -1)

	response := my_http.Response{
		Protocol:   "HTTP/1.1",
		StatusCode: 200,
		StatusMsg:  "OK",
		Body:       []byte(body),
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
		Body:       []byte(body),
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(body)),
		},
	}

	return response.WriteTo(c)
}

func GetFile(r *my_http.Request, c net.Conn, dir string) error {
	var fileFound bool
	var body []byte

	var err error
	if string(dir[len(dir)-1]) != "/" {
		dir = dir + "/"
	}

	directory, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal("Failed to open static files directory, ", err)
	}

	fileName := strings.Replace(r.Path, "/files/", "", -1)

	for _, v := range directory {
		if v.Name() == fileName {
			fileFound = true

			body, err = os.ReadFile(dir + v.Name())
			if err != nil {
				return err
			}
		}
	}

	if !fileFound {
		return NotFound(c)
	} else {
		response := my_http.Response{
			Protocol:   "HTTP/1.1",
			StatusCode: 200,
			StatusMsg:  "OK",
			Body:       body,
			Headers: map[string]string{
				"Content-Type":   "application/octet-stream",
				"Content-Length": fmt.Sprintf("%d", len(body)),
			},
		}

		return response.WriteTo(c)
	}
}

func PostFile(r *my_http.Request, c net.Conn, dir string) error {
	fileName := strings.Replace(r.Path, "/files/", "", -1)

	if string(dir[len(dir)-1]) != "/" {
		dir = dir + "/"
	}

	file, err := os.Create(dir + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.Write(r.Body); err != nil {
		log.Fatal(err)
	}

	file.Sync()

	return Created(c)
}
