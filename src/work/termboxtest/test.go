package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func keyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				fmt.Println(termbox.EventKey)
				termbox.Flush()
			}
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	keyEvent()
}
