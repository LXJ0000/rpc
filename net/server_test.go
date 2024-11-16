package net

import "testing"

func TestServe(t *testing.T) {
	s := NewServer()
	err := s.Start("tcp", ":8080")
	if err != nil {
		t.Error(err)
	}
}
