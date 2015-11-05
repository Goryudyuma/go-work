package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net"
	"net/url"
	"os"
	"strings"
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

		conn, err := listner.Accept()
		if err != nil {
			continue
		}
		if handleClient(conn) {
			continue
		}

		fp, err := os.Open("user.txt")
		if err != nil {
			panic(err)
		}

		s := bufio.NewScanner(fp)
		for s.Scan() {
			//api.PostDMToScreenName("召集がかかりました！ "+time.Now().String(), s.Text())
			//fmt.Println(M)
		}
		fp.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

type jsondata struct {
	ID   string `json:"id"`
	TEXT string `json:"text"`
}

func handleClient(conn net.Conn) bool {
	defer conn.Close()
	/*
		//conn.SetReadDeadline(time.Now().Add(100 * time.Second))
		//fmt.Println("client accept!")
		messageBuf := make([]byte, 1024)
		messageLen, _ := conn.Read(messageBuf)
		//checkError(err)

		message := string(messageBuf[:messageLen])
		//message = strings.Trim(message, "\n")
		//fmt.Println(message)
		mes := strings.Split(message, "keep-alive")
		fmt.Println(mes)

		fmt.Println(mes[1])
		message = mes[1]
		opt, _ := url.ParseQuery(message)
		fmt.Println("!")
		fmt.Println(opt["token"])

		fmt.Println("!")

		var data interface{}
		dec := json.NewDecoder(strings.NewReader(message))
		dec.Decode(&data)
		//fmt.Println(data)
	*/
	return true
}
