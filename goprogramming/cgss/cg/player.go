package cg

import "fmt"

/**
为了便于演示聊天系统，我们为每个玩家都起了一个独立的goroutine，监听所有发送给他们的聊天信息，
一旦收到就即时打印到控制台上。
*/

type Player struct {
	Name  string
	Level int
	Exp   int
	Room  int
	mq    chan *Message
}

func Newplayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{
		Name:  "",
		Level: 0,
		Exp:   0,
		Room:  0,
		mq:    m,
	}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "收到消息：", msg.Content)
		}
	}(player)

	return player
}
