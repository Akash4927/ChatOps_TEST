package main

import (
	"fmt"
	"net"
	"testing"
)

func Test_main(t *testing.T) {
	server, client := net.Pipe()
	go func() {
		conn, err := net.Dial("tcp", ":9090")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := fmt.Fprintln(conn); err != nil {
			t.Fatal(err)
		}
		server.Close()
	}()

	_, err := net.Listen("tcp", ":9090")
	if err != nil {
		t.Fatal(err)
	}
	client.Close()
}
