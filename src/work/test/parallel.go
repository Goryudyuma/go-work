package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	// 1秒かかるコマンド
	log.Print("sleep1 started.")
	time.Sleep(1 * time.Second)
	log.Print("sleep1 finished.")

	// 2秒かかるコマンド
	log.Print("sleep2 started.")
	time.Sleep(2 * time.Second)
	log.Print("sleep2 finished.")

	// 3秒かかるコマンド
	log.Print("sleep3 started.")
	time.Sleep(3 * time.Second)
	log.Print("sleep3 finished.")

	log.Print("all finished.")
}
