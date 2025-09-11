package main

import (
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"office-helper/examples/websocket/client"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	chatGpt = client.NewChat()
)

func server(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		s, err := chatGpt.Message(context.Background(), string(message))
		if err != nil {
			break
		}
		if err = c.WriteMessage(mt, []byte(s)); err != nil {
			break
		}

	}

}

func main() {
	http.HandleFunc("/ws", server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
