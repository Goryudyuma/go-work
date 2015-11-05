package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started.")

	sleep1_finished := make(chan bool)
	sleep2_finished := make(chan bool)
	sleep3_finished := make(chan bool)

	go func() {
		log.Print("sleep1 started.")
		time.Sleep(1 * time.Second)
		log.Print("sleep1 finished.")
		sleep1_finished <- true
	}()

	go func() {
		log.Print("sleep2 started.")
		time.Sleep(2 * time.Second)
		log.Print("sleep2 finished.")
		sleep2_finished <- true
	}()

	go func() {
		log.Print("sleep3 started.")
		time.Sleep(3 * time.Second)
		log.Print("sleep3 finished.")
		sleep3_finished <- true
	}()

	<-sleep1_finished
	<-sleep2_finished
	<-sleep3_finished

	log.Print("all finished.")

}
