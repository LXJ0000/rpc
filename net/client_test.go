package net

import (
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	c := NewClient("tcp", "localhost:8080", time.Second*3)
	resp, err := c.Send("hello")
	if err != nil {
		t.Error(err)
	}
	t.Error(resp)
}
