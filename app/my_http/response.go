package my_http

import (
	"bytes"
	"fmt"
	"net"
)

type Response struct {
	Protocol   string
	StatusCode int
	StatusMsg  string
	Headers    map[string]string
	Body       []byte
}

func (r *Response) WriteTo(conn net.Conn) error {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s %d %s\r\n", r.Protocol, r.StatusCode, r.StatusMsg))
	for header, data := range r.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", header, data))
	}
	buf.WriteString("\r\n")
	buf.Write(r.Body)

	if _, err := conn.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}
