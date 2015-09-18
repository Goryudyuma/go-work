package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				break
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		i, j := termbox.Size()
		for x := 0; x < len(strconv.Itoa(i)); x++ {
			termbox.SetCell(x, 0, []rune(strconv.Itoa(i))[x], 0, 0)
		}
		for x := 0; x < len(strconv.Itoa(j)); x++ {
			termbox.SetCell(x, 2, []rune(strconv.Itoa(j))[x], 0, 0)
		}
		termbox.Flush()
	}
}
