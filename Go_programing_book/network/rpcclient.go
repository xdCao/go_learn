package main

import (
	"fmt"
	"net/rpc"
	myrpc "xdCao/golearn/goprogramming/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("connect failed", err)
		return
	}
	var reply int
	args := myrpc.Args{7, 8}
	// 同步调用
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("arith error, ", err)
		return
	}
	fmt.Println("Arith called, %d * %d = %d", args.A, args.B, reply)

	quotient := new(myrpc.Quotient)
	// 异步调用
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done
	fmt.Printf("replyCall: %v\n", replyCall)
}
