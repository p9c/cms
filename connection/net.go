package connection

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	Socket net.Conn
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

func (manager *ClientManager) Receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.Socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

func (client *Client) Receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.Socket.Read(message)
		if err != nil {
			client.Socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}

func (manager *ClientManager) Send(client *Client) {
	defer client.Socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.Socket.Write(message)
		}
	}
}

func StartServerModule() {
	fmt.Println("Starting server...")
	listener, error := net.Listen("tcp", ":19999")
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
		client := &Client{Socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.Receive(client)
		go manager.Send(client)
	}
}

func StartClientModule(c net.Conn) (net.Conn, error) {
	fmt.Println("Starting client...")
	c, err := net.Dial("tcp", "localhost:19999")
	if err != nil {
		fmt.Println(err)
	}
	client := &Client{Socket: c}
	go client.Receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		c.Write([]byte(strings.TrimRight(message, "\n")))
	}
	return c, err
}
