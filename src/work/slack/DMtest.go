package main

import (
	"bufio"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net"
	"os"
	"time"
)

func access(fp *os.File) *anaconda.TwitterApi {
	scanner := bufio.NewScanner(fp)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	scanner.Scan()
	consumer_key := scanner.Text()
	scanner.Scan()
	consumer_secret := scanner.Text()
	scanner.Scan()
	accesstoken := scanner.Text()
	scanner.Scan()
	accesstoken_secret := scanner.Text()

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	return anaconda.NewTwitterApi(accesstoken, accesstoken_secret)
}

var api *anaconda.TwitterApi

func main() {
	api_chan := make(chan *anaconda.TwitterApi)

	go func() {
		fp, err := os.Open("/etc/pass/twi_sl.txt")
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		api_chan <- access(fp)
	}()

	service := ":55556"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	api = <-api_chan
	for {

		_, err := listner.Accept()
		if err != nil {
			continue
		}

		fp, err := os.Open("user.txt")
		if err != nil {
			panic(err)
		}

		s := bufio.NewScanner(fp)
		for s.Scan() {
			api.PostDMToScreenName("召集がかかりました！ "+time.Now().String(), s.Text())
			//fmt.Println(M)
		}
		//go handleClient(conn)
		fp.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

/*
func handleClient(conn net.Conn) {
	defer conn.Close()
	//conn.SetReadDeadline(time.Now().Add(100 * time.Second))
	//fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, _ := conn.Read(messageBuf)
	//checkError(err)

	message := string(messageBuf[:messageLen])
	//message = strings.Trim(message, "\n")
	//fmt.Println(message)
}
*/
