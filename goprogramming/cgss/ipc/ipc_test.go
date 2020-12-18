package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handler(method, params string) *Response {
	return &Response{
		Code: "OK",
		Body: "ECHO" + method + "~" + params,
	}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("foo", "from Client1")
	resp2, _ := client2.Call("foo", "from Client2")

	fmt.Println(resp1)
	fmt.Println(resp2)

	client1.Close()
	client2.Close()

}
