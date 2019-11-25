package main

import (
	"net/rpc"
	"log"
	"fmt"
)

// 服务名字
const HelloServiceName = "path/to/pkg.HelloService"

// 服务详细方法列表
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 注册该类服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
