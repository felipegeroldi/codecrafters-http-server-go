package my_http

import "strings"

const (
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Headers  map[string]string
}

func ParseData(data []byte) *Request {
	dataStrs := strings.Split(string(data), "\r\n")

	reqData := strings.Split(dataStrs[0], " ")
	dataStrs = dataStrs[1:]

	req := &Request{
		Method:   reqData[0],
		Path:     reqData[1],
		Protocol: reqData[2],
		Headers:  make(map[string]string, 0),
	}

	for _, headerData := range dataStrs {
		header := strings.Split(headerData, ": ")
		if len(header) > 1 {
			req.Headers[header[0]] = header[1]
		}
	}

	return req
}
