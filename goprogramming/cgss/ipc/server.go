package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handler(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c

			//关闭该链接
			if request == "close" {
				fmt.Println("close session ")
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("错误的requst格式:", request)
				return
			}

			fmt.Println("server recive :", req)
			resp := server.Handler(req.Method, req.Params)
			b, err := json.Marshal(resp)
			if err != nil {
				fmt.Println("marshal json 失败:", resp)
				return
			}
			//返回结果
			c <- string(b)
		}
	}(session)

	fmt.Println("新的session创建成功")
	return session
}
