package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) recieve(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECIEVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

func startServerMode(addr string) {
	fmt.Println("Starting server...")
	listener, error := net.Listen("tcp", addr)
	if error != nil {
		fmt.Println(error)
	}
	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.recieve(client)
		go manager.send(client)
	}
}

//Now we shift focus to the client side

func (client *Client) recieve() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECIEVED: " + string(message))
		}
	}
}

func startClientMode(addr string) {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", addr)
	if error != nil {
		fmt.Println(error)
	}
	client := &Client{socket: connection}
	go client.recieve()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		if string(message) == "end\n" {
			terminationMessage := "Client user has terminated the session"
			connection.Write([]byte(strings.TrimRight(terminationMessage, "\n")))
			fmt.Println("Terminating client")
			break
		}
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
}

func fetchAddr() string {
	cmd := "hostname"
	args := []string{"-I"}
	command := exec.Command(cmd, args[0])
	out, err := command.Output()

	if err != nil {
		fmt.Println("Error", err)
		panic(err)
		return ""
	}

	port := "42070"
	return strings.TrimSpace(string(out)) + ":" + port
}

func main() {
	addr := fetchAddr()
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServerMode(addr)
	} else {
		serverAddr := addr //This is just for local testing
		startClientMode(serverAddr)
	}

}
