package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var mutex = &sync.Mutex{}
var isWsConnected = false

func WSHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	isWsConnected = true
	mutex.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
	}
}

func IsWsConnected() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return isWsConnected
}

func TriggerHandler(c *gin.Context) {
	broadcast <- "refresh"
}

func TriggerRefresh() {
	broadcast <- "refresh"
	log.Println("Trigger Refresh called")
}

func init() {
	go func() {
		for {
			msg := <-broadcast
			mutex.Lock()
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Printf("Websocket error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
			mutex.Unlock()
		}
	}()
}
