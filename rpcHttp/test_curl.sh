curl localhost:1234/jsonrpc -X POST \
    --data '{"method":"path/to/pkg.HelloService.Hello","params":["world"],"id":0}'