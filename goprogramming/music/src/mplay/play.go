package mplay

import (
	"fmt"
	"strings"
)

type Player interface {
	Play(source string, channel chan bool)
}

func Play(source, mtype string, channel chan bool) {
	var p Player
	mtype = strings.ToLower(mtype)
	switch mtype {
	case "mp3":
		p = &Mp3Player{}
	case "wav":
		p = &WavPlayer{}
	default:
		fmt.Println("不支持的格式: ", mtype)
	}

	p.Play(source, channel)
}
