//http://kudohamu.hatenablog.com/entry/2014/11/03/071802
package main

import (
	"bytes"
	"net"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	service := ":55555"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(100 * time.Second))
	//fmt.Println("client accept!")
	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	checkError(err)

	message := string(messageBuf[:messageLen])
	//message = strings.Trim(message, "\n")
	//fmt.Println(message)
	if check_regexp(`[^(KILHZON\+\-\.\d\s)]`, message) {
		//fmt.Println(message)
		out := []byte("error2\n")
		//fmt.Println("error2")
		conn.SetWriteDeadline(time.Now().Add(100 * time.Second))
		conn.Write([]byte(out))
	} else {
		cmd := exec.Command("/root/work/cpp/dobutsu/checkState", message)
		if err != nil {
			//fmt.Println(err)
			//os.Exit(1)
		}

		out, err := cmd.Output()
		if err != nil {
			//fmt.Println(err)
			out = []byte("error\n")
			//os.Exit(1)
		} else {
			out2 := bytes.SplitAfter(out, []byte("Move : "))
			out3 := bytes.SplitAfter(out2[1], []byte("\n"))
			out = out3[0]
		}

		//fmt.Println(err)
		//fmt.Println(out)
		//out = strings.Trim(out, "\n")
		conn.SetWriteDeadline(time.Now().Add(100 * time.Second))
		conn.Write([]byte(out))
	}
}

func checkError(err error) {
	if err != nil {
		//fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		//os.Exit(1)
	}
}
