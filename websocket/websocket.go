package websocketHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	youtubestats "youtube-subscription/youtube"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WebSocket is a type of connection struct
type WebSocket struct {
	Conn *websocket.Conn
}

// NewWebSocketConn creates and upgrade websocket protocol from http
// Create new goroutine to update the value constantly
func NewWebSocketConn(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("Error happend when updating from http to websocket: ", err)
		return nil, err
	}
	// defer conn.Close()
	// check current header info
	header := r.Header
	fmt.Println("new websocket current header: ", header)
	// add connection value to struct
	ws := &WebSocket{
		Conn: conn,
	}
	// start new goroutine with WebSocket struct that updates youtube subscription number
	go ws.Update()
	return ws, nil
}

// Update : Constantly update the YouTube data value coming from GetYoutubeData
// function as struct and convert it into JSON format so we can
// write out our message
func (ws *WebSocket) Update() {
	// create forever for loop so we can constantly update the value
	for {
		updateTicker := time.NewTicker(3 * time.Second)

		for t := range updateTicker.C {
			fmt.Printf("Just Updated Youtube Data: %+v\n", t)

			items, err := youtubestats.GetYoutubeData()
			if err != nil {
				fmt.Println(err)
			}
			// items is stored as struct so use marshal to convert into JSON
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println("Marshal failed: ", err)
			}
			// if the method returns something, it should close the for loop
			err = ws.Conn.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil {
				fmt.Println("Error happend when writting message: ", err)
				break // close connection
			}
		}
	}
}
