package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a port number: ")
	portInput, _ := reader.ReadString('\n')
	portInput = portInput[:len(portInput)-1]

	port, err := strconv.Atoi(strings.TrimSpace(portInput))
	if err != nil {
		fmt.Println("Invalid port number:", err)
		return
	}

	log.Printf("Starting server on port: %d \n", port)
	http.HandleFunc("/ws", handleWebSocket)

	err = http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		log.Println(err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	log.Printf("new connection! %s \n", conn.RemoteAddr().String())
	if err != nil {
		log.Println(err)
		return
	}
	defer closeConnection(conn)

	// Main logic here.
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var receivedCoordinate POINT

		switch strMsg := string(msg[:]); strMsg {
		case "LeftClick":
			LeftClick()
		default:
			json.Unmarshal([]byte(msg), &receivedCoordinate)
			SetCursorPosition(receivedCoordinate.X, receivedCoordinate.Y)
		}

	}
}

func closeConnection(conn *websocket.Conn) {
	fmt.Printf("Connection closed by %s \n", conn.RemoteAddr().String())
	conn.Close()
}
