package net

import (
	"net"
)

const (
	lenOfMessage = 8
)

func Serve(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func() {
			if err := handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}

func handleConn(conn net.Conn) error {
	// 1 读数据
	// 2 处理数据
	// 3 返回响应
	for {
		byteDatas := make([]byte, lenOfMessage)
		if _, err := conn.Read(byteDatas); err != nil {
			return err
		}
		resp := handleMeaasge(byteDatas)
		if _, err := conn.Write(resp); err != nil {
			return err
		}
	}
}

func handleMeaasge(data []byte) []byte {
	res := make([]byte, len(data)*2)
	copy(res[:len(data)], data)
	copy(res[len(data):], data)
	return res
}
