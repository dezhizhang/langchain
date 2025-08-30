package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"office-helper/openai"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	gpt = openai.NewChat()
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
		s, err := gpt.Message(context.Background(), string(message))
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
	fmt.Println("websocket server start")
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		panic(err)
	}
}
