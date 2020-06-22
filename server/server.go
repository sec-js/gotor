package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KingAkeem/goTor/server/gobot"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, _ := wsUpgrader.Upgrade(w, r, nil)
	defer wsConn.Close()
	for {
		msg := struct {
			Type string `json:"type"`
			Link string `json:"link"`
		}{}
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Unable to decode WebSocket message. Error: %+v", err)
			return
		}
		linkChan, err := gobot.GetLinks(msg.Link)
		if err != nil {
			log.Printf("Unable to get links. Error: %+v", err)
			return
		}
		for link := range linkChan {
			err := wsConn.WriteJSON(struct {
				Type     string     `json:"type"`
				LinkData gobot.Link `json:"linkData"`
			}{
				Type:     "GET_LINK_RESULT",
				LinkData: link,
			})
			if err != nil {
				log.Printf("Unable to write to WebSocket connection. Error: %v", err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", getLinksHandler)
	fmt.Println("Serving on localhost:3050")
	log.Fatal(http.ListenAndServe(":3050", nil))
}
