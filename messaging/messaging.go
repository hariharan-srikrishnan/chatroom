package messaging

import (
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	"io"
	"bytes"
	"chatroom/utils"
)

type void struct{}

var connectedClients map[*websocket.Conn]void
var messageHistory []string

func init() {
	connectedClients = make(map[*websocket.Conn]void)
	messageHistory = make([]string, 0)
}


func ChatServer(conn *websocket.Conn) {
	fmt.Println("Received request for path /chat")
	connectedClients[conn] = void{}
	defer func() {
		delete(connectedClients, conn)
		conn.Close()
	}()

	for {
		var message string
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Printf("connection closed by client %+v\n", conn)
			default:
				fmt.Println("error in receiving a message", err)
			}
			break
		}
		fmt.Println("data: ", message)
		messageHistory = append(messageHistory, message)
		fmt.Println("message history: ", messageHistory)
		broadcastMessage(utils.StrToBytes(&message))
	}

}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for path /check")
	w.Write(utils.OK_BYTES)
}

func MessageHistoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for path /messageHistory")
	fmt.Println("message history: ", messageHistory)
	fmt.Println("connected clients: ", connectedClients)

	response := []byte{}
	dataBuffer := bytes.NewBuffer(response)
	for _, message := range messageHistory {
		dataBuffer.WriteString(message)
		dataBuffer.WriteRune('\n')
	}
	fmt.Println(dataBuffer.Bytes())
	w.Write(dataBuffer.Bytes())
}

func broadcastMessage(message []byte) {
	for client := range connectedClients {
		// go func(client *websocket.Conn) {
		// sendMessage(client, message, &wg)
		// }(client)
		sendMessage(client, message)
	}
}

func sendMessage(client *websocket.Conn, message []byte) {
	err := websocket.Message.Send(client, string(message))
	if err != nil {
		_ = fmt.Errorf("unable to send a message, error: %v", err)
	}
}
