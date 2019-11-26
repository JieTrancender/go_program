package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
	"go_program/rpcJson/rpcJson"
)

func main() {
	conn, err:= net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial wrong:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(rpcJson.HelloServiceName+".Hello", "world", &reply)
	if err != nil {
		log.Fatal("call wrong:", err)
	}

	fmt.Println(reply)
}
