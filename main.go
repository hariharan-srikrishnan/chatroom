package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"chatroom/messaging"
)

func main() {
	http.HandleFunc("/check", messaging.CheckHandler)

	http.HandleFunc("/chat",
		func(w http.ResponseWriter, r *http.Request) {
			// creating our own server to bypass origin check failures
			wsServer := websocket.Server{Handler: websocket.Handler(messaging.ChatServer)}
			wsServer.ServeHTTP(w, r)
		},
	)

	http.HandleFunc("/messageHistory", messaging.MessageHistoryHandler)

	fmt.Println("starting http server...")
	done := make(chan int)
	go func() {
		err := http.ListenAndServe(":15321", nil)
		if err != nil {
			fmt.Print("Unable to start http server", err)
			panic(err)
		}
		done <- 0
	}()
	fmt.Println("http server started")
	<-done
}
