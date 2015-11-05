//http://kudohamu.hatenablog.com/entry/2014/11/03/071802
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s message", os.Args[0])
		os.Exit(1)
	}
	message := os.Args[1]

	serverIP := "127.0.0.1" //サーバ側のIP
	serverPort := "55555"   //サーバ側のポート番号
	myIP := "127.0.0.1"     //クライアント側のIP
	myPort := 55556         //クライアント側のポート番号

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	checkError(err)
	myAddr := new(net.TCPAddr)
	myAddr.IP = net.ParseIP(myIP)
	myAddr.Port = myPort
	conn, err := net.DialTCP("tcp", myAddr, tcpAddr)
	checkError(err)
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))

	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readlen, err := conn.Read(readBuf)
	checkError(err)

	fmt.Println("server: " + string(readBuf[:readlen]))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
