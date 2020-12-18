package mplay

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"time"
)

type WavPlayer struct {
	stat     int
	progress int
}

func (p WavPlayer) Play(source string, channel chan bool) {
	f, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	defer speaker.Close()
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	for {
		select {
		case <-done:
			fmt.Println("play over")
			return
		case value := <-channel:
			fmt.Printf("value %t %v\n", value, value)
			return
		default:
			//fmt.Println("playing ", source)
			//time.Sleep(time.Second)
		}
	}
}
