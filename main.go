package main

import (
	"fmt"
	"log"
	"net/http"

	websocketHandler "youtube-subscription/websocket"
	youtubestats "youtube-subscription/youtube"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocketHandler.NewWebSocketConn(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		panic(err)
	}
	fmt.Println(ws)
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	item, err := youtubestats.GetYoutubeData()
	if err != nil {
		log.Fatalln("Error happend when loading youtube data: ", err)
	}
	fmt.Printf("%+v\n", item)

	setupRoutes()
}
