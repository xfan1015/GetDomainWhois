package models

import (
	"bytes"
	// "fmt"
	"io"
	"net"
)

//得到域名的whois完整信息
func GetDomainWhois(service, domain string) (string, error) {
	service = service + ":43"
	conn, err := net.Dial("tcp", service)
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", err
	}

	result, err := readFully(conn)
	if err != nil {
		return "", err
	}
	// fmt.Println(result)
	return string(result), nil

}

//readFully完整的读取whois信息，并返回完整结果
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close() //关闭连接
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
