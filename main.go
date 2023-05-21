package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Main logic here.
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var receivedCoordinate POINT
		json.Unmarshal([]byte(msg), &receivedCoordinate)
		log.Printf("received X: %d received Y: %d", receivedCoordinate.X, receivedCoordinate.Y)
		SetCursorPosition(receivedCoordinate.X, receivedCoordinate.Y)
	}
}
