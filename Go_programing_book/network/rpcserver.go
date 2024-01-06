package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
	myrpc "xdCao/golearn/goprogramming/rpc"
)

func main() {
	arith := new(myrpc.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:9999")
	if e != nil {
		fmt.Println("listen error", e)
		return
	}
	go http.Serve(l, nil)
	for {
		time.Sleep(1 * time.Microsecond)
	}
}
