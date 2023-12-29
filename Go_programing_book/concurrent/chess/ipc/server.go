package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{Server: server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string)

	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Printf("Invalid request: %v\n", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, _ := json.Marshal(resp)
			c <- string(b)
		}
		fmt.Println("session closed")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}
