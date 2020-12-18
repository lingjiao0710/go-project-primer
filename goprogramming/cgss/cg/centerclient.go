package cg

import (
	"encoding/json"
	"errors"

	"github.com/lingjiao0710/test/goprogramming/cgss/ipc"
)

/**
CenterClient匿名组合了IpcClient,
这样就可以直接在代码中调用IpcClient的功能了。
*/

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}

	return err
}

func (client *CenterClient) RemovePlayer(name string) error {

	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message} //构造Message结构体

	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
