package net

import (
	"fmt"
	"net"
	"time"
)

func Connect(network, addr string) error {
	conn, err := net.DialTimeout(network, addr, time.Second*3)
	if err != nil {
		return err
	}
	for {
		data := "hello"
		byteDatas := []byte(data)
		if _, err := conn.Write(byteDatas); err != nil {
			return err
		}

		resp := make([]byte, 1024)
		if _, err := conn.Read(resp); err != nil {
			return err
		}
		fmt.Println(string(resp))

		return nil
	}
}
