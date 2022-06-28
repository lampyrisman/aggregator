package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"skv"
	s "structs"
)

const (
	IP   = "127.0.0.1"
	PORT = "10000"
	TYPE = "tcp"
)

var services = []string{"r_queue", "render", "p_queue", "parser", "f_queue", "formatter"}

func cerr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	channels := make(map[string]chan s.Reg_data)
	for _, v := range services {
		channels[v] = make(chan s.Reg_data)
		go skv.Reg(channels[v])
	}

	var err error
	listen, err := net.Listen(TYPE, IP+":"+PORT)
	cerr(err)

	for {
		var inData s.Reg_data
		conn, _ := listen.Accept()
		inMsg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("recv", inMsg)
		err = json.Unmarshal([]byte(inMsg), &inData)
		channels[inData.Role] <- inData
		outData := <-channels[inData.Role]
		fmt.Println("main got", outData, "\n")
		out_bytes, err := json.Marshal(outData)
		cerr(err)
		conn.Write(out_bytes)
		conn.Close()
	}
}
