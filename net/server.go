package net

import (
	"encoding/binary"
	"net"
)

const (
	lenOfHeader = 8 // 8 bit 64 位
)

func (s *Server) handleConn(conn net.Conn) error {
	// 1 读数据 (长度 + 数据)
	// 2 处理数据
	// 3 返回响应
	for {
		// req = header + reqData
		header := make([]byte, lenOfHeader) // 8字节长度 且 代表数据长度
		if _, err := conn.Read(header); err != nil {
			return err
		}
		length := binary.BigEndian.Uint64(header)
		reqData := make([]byte, length)
		if _, err := conn.Read(reqData); err != nil {
			return err
		}

		// resp = header + respData
		respData := s.handleMessage(reqData)
		resp := make([]byte, lenOfHeader+len(respData))
		binary.BigEndian.PutUint64(resp[:lenOfHeader], uint64(len(respData)))
		copy(resp[lenOfHeader:], respData)

		if _, err := conn.Write(resp); err != nil {
			return err
		}
	}
}

func (s *Server) handleMessage(data []byte) []byte {
	res := make([]byte, len(data)*2)
	copy(res[:len(data)], data)
	copy(res[len(data):], data)
	return res
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(network, addr string) error {
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
			if err := s.handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}
