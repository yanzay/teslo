package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo/templates"
)

type Message struct {
	ID      string
	Content string
}

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WSHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.WriteHello(w, "%username%")
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for {
		writer, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Sending %d", i)
		templates.WriteMessage(writer, "app", templates.Body(fmt.Sprint(i)))
		err = writer.Close()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
		i++
	}
}
