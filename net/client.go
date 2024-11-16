package net

import (
	"encoding/binary"
	"net"
	"time"
)

type Client struct {
	network string
	addr    string
	timeout time.Duration
}

func NewClient(network, addr string, timeout time.Duration) *Client {
	return &Client{
		network: network,
		addr:    addr,
		timeout: timeout,
	}
}

func (c *Client) Send(data string) (string, error) {
	conn, err := net.DialTimeout(c.network, c.addr, c.timeout)
	if err != nil {
		return "", err
	}
	for {
		req := make([]byte, lenOfHeader+len(data))
		binary.BigEndian.PutUint64(req[:lenOfHeader], uint64(len(data)))
		copy(req[lenOfHeader:], data)
		if _, err := conn.Write(req); err != nil {
			return "", err
		}

		header := make([]byte, lenOfHeader) // 8字节长度 且 代表数据长度
		if _, err := conn.Read(header); err != nil {
			return "", err
		}
		length := binary.BigEndian.Uint64(header)
		respData := make([]byte, length)
		if _, err := conn.Read(respData); err != nil {
			return "", err
		}

		return string(respData), nil
	}
}
