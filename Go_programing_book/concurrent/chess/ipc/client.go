package ipc

import "encoding/json"

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{conn: c}
}

func (client *IpcClient) Call(method, params string) *Response {
	req := &Request{method, params}
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	client.conn <- string(b)
	respStr := <-client.conn
	resp := &Response{}
	err = json.Unmarshal([]byte(respStr), resp)
	if err != nil {
		return nil
	}
	return resp
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
