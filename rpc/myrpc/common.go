package myrpc

const HelloServiceName = "path/to/pkg.HelloService"

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
