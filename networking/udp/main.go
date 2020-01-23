package main

import (
	"fmt"
	"net"
)

type Client struct {
	id     int
	socket net.Conn
	data   chan []byte
}

func (client *Client) recieve() {
	for {
		buffer := make([]byte, 1024)
		length, _ := client.socket.Read(buffer)
		if length > 0 {
			fmt.Println("RECIEVED: " + string(buffer))
		}
	}
}

func (client *Client) send() {
	client.socket.Write([]byte("Hello server, this is client " + string(client.id)))
}

func startClientMode(id int) {
	fmt.Println("Starting client with id " + string(id))
	connection, err := net.Dial("udp", "localhost:12345")
	if err != nil {
		fmt.Println(err)
	}
	client := &Client{
		id:     id,
		socket: connection,
	}
	go client.recieve()
}

func startServerMode() {
	fmt.Println("Starting server...")
	pc, err := net.ListenPacket("udp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	defer pc.Close
}

//http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/