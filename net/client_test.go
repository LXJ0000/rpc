package net

import (
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	go func() {
		if err := Serve("tcp", ":8080"); err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second * 3)
	if err := Connect("tcp", "localhost:8080"); err != nil {
		t.Error(err)
	}
}
