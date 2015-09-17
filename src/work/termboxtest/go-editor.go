package main

import (
	"bufio"
	"flag"
	"github.com/nsf/termbox-go"
	"math"
	"os"
	"strconv"
)

var nowline, nowcol int
var runes [][]rune
var filename string

func draw() {
	linecount := 3
	for maxline := len(runes) - 1; maxline != 0; maxline /= 10 {
		linecount++
	}
	for i := range runes {
		j := 0
		if i == nowline {
			j = nowline + 1
		} else {
			j = int(math.Abs(float64(nowline - i)))
		}
		linerune := []rune(strconv.Itoa(j))
		for j := range linerune {
			termbox.SetCell(j, i, linerune[j], 0, 0)
		}
		for j := range runes[i] {
			if i == nowline && j == nowcol {
				termbox.SetCell(j+linecount, i, runes[i][j], 1, termbox.Attribute(5)+1)
			} else {
				termbox.SetCell(j+linecount, i, runes[i][j], 0, 0)
			}
		}
	}
	if len(runes[nowline]) <= nowcol {
		termbox.SetCell(nowcol+linecount, nowline, ' ', 0, termbox.Attribute(5)+1)
	}
	termbox.Flush()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func keyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				{
					if nowcol > 0 {
						nowcol--
						runes[nowline] = append(runes[nowline][:nowcol], runes[nowline][nowcol+1:]...)
					}
				}
			case termbox.KeyEnter:
				{
					nowline++
					s := make([][]rune, len(runes)+1)
					copy(s, runes[:nowline])
					s[nowline] = make([]rune, 0)
					copy(s[nowline+1:], runes[nowline:])
					runes = s
					if nowcol > len(runes[nowline]) {
						nowcol = len(runes[nowline])
					}
				}
			case termbox.KeyArrowLeft:
				{
					nowcol--
					if nowcol < 0 {
						nowcol = 0
					}
				}
			case termbox.KeyArrowRight:
				{
					nowcol++
					if nowcol > len(runes[nowline]) {
						nowcol = len(runes[nowline])
					}
				}
			case termbox.KeyArrowUp:
				{
					nowline--
					if nowline < 0 {
						nowline = 0
					}

					if nowcol > len(runes[nowline]) {
						nowcol = len(runes[nowline])
					}
				}
			case termbox.KeyArrowDown:
				{
					nowline++
					if nowline > len(runes)-1 {
						nowline = len(runes) - 1
					}
					if nowcol > len(runes[nowline]) {
						nowcol = len(runes[nowline])
					}
				}
			default:
				s := make([]rune, len(runes[nowline])+1)
				copy(s, runes[nowline][:nowcol])
				s[nowcol] = ev.Ch
				copy(s[nowcol+1:], runes[nowline][nowcol:])
				runes[nowline] = s
				nowcol++
			}
			draw()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func Exists() bool {
	_, err := os.Stat(filename)
	return err == nil
}

func construct(fp *os.File) {
	runes = make([][]rune, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() || len(runes) == 0 {
		runes = append(runes, []rune(scanner.Text()))
	}
	nowline = 0
	nowcol = 0
	draw()
}

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		filename = flag.Args()[0]
	} else {
		i := 0
		for {
			filename = strconv.Itoa(i) + ".txt"
			if !Exists() {
				break
			}
			i++
		}

	}
	var fp *os.File
	var err error
	if Exists() {
		fp, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
	} else {
		fp, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	defer fp.Close()

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	construct(fp)
	keyEvent()
}
