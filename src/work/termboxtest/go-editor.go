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
	W, H := termbox.Size()
	var DisplayX, DisplayY int

	if nowcol-W/2 < 0 {
		DisplayX = 0
	} else {
		DisplayX = nowcol - W/2
	}

	if nowline-H/2 < 0 {
		DisplayY = 0
	} else {
		DisplayY = nowline - H/2
	}

	for maxline := len(runes) - 1; maxline != 0; maxline /= 10 {
		linecount++
	}

	for k := 0; k < H && DisplayY+k < len(runes); k++ {
		i := DisplayY + k + DisplayX*0
		j := 0
		if i == nowline {
			j = nowline + 1
		} else {
			j = int(math.Abs(float64(nowline - i)))
		}
		linerune := []rune(strconv.Itoa(j))
		for l := range linerune {
			termbox.SetCell(l, k, linerune[l], 0, 0)
		}
		for l := 0; l < W && l+DisplayX < len(runes[i]); l++ {
			if i == nowline && l+DisplayX == nowcol {
				termbox.SetCell(l+linecount, k, runes[i][l+DisplayX], 1, termbox.Attribute(5)+1)
			} else {
				termbox.SetCell(l+linecount, k, runes[i][l+DisplayX], 0, 0)
			}
		}
	}
	if len(runes[nowline]) <= nowcol {
		termbox.SetCell(nowcol+linecount-DisplayX, nowline-DisplayY, ' ', 0, termbox.Attribute(5)+1)
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
					copy(s, runes[:nowline-1])
					s[nowline-1] = make([]rune, len(runes[nowline-1][:nowcol]))
					s[nowline] = make([]rune, len(runes[nowline-1][nowcol:]))
					copy(s[nowline-1], runes[nowline-1][:nowcol])
					copy(s[nowline], runes[nowline-1][nowcol:])
					copy(s[nowline+1:], runes[nowline:])
					runes = s
					nowcol = 0
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
