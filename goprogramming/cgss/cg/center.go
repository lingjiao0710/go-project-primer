package cg

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/lingjiao0710/test/goprogramming/cgss/ipc"
)

/**
我们为中央服务器实现了几个示范用的指令：
添加用户、删除用户、列出用户和广播。
为了便于调用这个服务器的功能，还需写一个centerclient.go
*/

var _ ipc.Server = &CenterServer{} //确认实现了Server接口

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}

func (server *CenterServer) Handler(method, params string) *ipc.Response {
	fmt.Println("handler :", method, params)
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
		return &ipc.Response{
			"200",
			players,
		}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
		return &ipc.Response{
			Code: "200",
		}
	default:
		return &ipc.Response{
			Code: "404",
			Body: method + " " + params,
		}
	}

	return &ipc.Response{
		Code: "200",
	}
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{
		servers: servers,
		players: players,
	}
}

func (server *CenterServer) addPlayer(params string) error {
	player := Newplayer()

	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	//没有做重复登录检查
	server.players = append(server.players, player)

	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[:i+1]...)
			}
			fmt.Println(server.players)
			return nil
		}
	}
	fmt.Println(server.players)
	return errors.New("找不到Player")
}

func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("listPlayer 没有Player在线")
	}
	return
}

func (server *CenterServer) broadcast(params string) error {
	var message Message

	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("broadcast 没有Player在线")
	}

	return err
}
