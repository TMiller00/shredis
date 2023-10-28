package main

import (
	"net"
	"testing"
)

func TestWrite(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	writer := NewWriter(serverConn)

	go func() {
		err := writer.Write(Value{typ: "string", str: ""})
		if err != nil {
			t.Errorf("Unexpected error. Got %v, expected nil", err)
		}
	}()

	err := serverConn.Close()
	if err != nil {
		t.Errorf("Unexpected error. Got %v, expected nil", err)
	}

	err = clientConn.Close()
	if err != nil {
		t.Errorf("Unexpected error. Got %v, expected nil", err)
	}
}
