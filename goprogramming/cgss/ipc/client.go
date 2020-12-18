package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	//c := make(chan string)
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}

	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn //等待返回值
	fmt.Println(str)

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	if err != nil {
		fmt.Println(err)
	}

	//var resp1 Request
	//err = json.Unmarshal([]byte(str), &resp1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp1)

	//resp = &Response{"123", "111"}
	resp = &resp1
	return
}

func (client *IpcClient) Close() {
	client.conn <- "close"
}
