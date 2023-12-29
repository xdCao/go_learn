package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server EchoServer) Name() string {
	return "EchoServer"
}

func (server EchoServer) Handle(method, params string) *Response {
	return &Response{Code: "200", Body: "ECHO: " + method + params}
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1 := client1.Call("From Client1", "")
	resp2 := client2.Call("From Client2", "")

	fmt.Printf("resp1: %v\n", resp1)
	fmt.Printf("resp2: %v\n", resp2)

	client1.Close()
	client2.Close()
}
