package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo/templates"
)

type Message struct {
	ID      string
	Content string
}

type Server struct {
	requests  chan Event
	responses chan Message
}

func NewServer() *Server {
	return &Server{
		requests:  make(chan Event, 0),
		responses: make(chan Message, 0),
	}
}

func (s *Server) Start() {
	go s.eventHandler()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", s.WSHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) eventHandler() {
	for {
		request := <-s.requests
		log.Println(request)
	}
}

var upgrader = websocket.Upgrader{}

func main() {
	server := NewServer()
	server.Start()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	items := []templates.Item{
		{ID: "1", Content: "write a todo app"},
		{ID: "2", Content: "extract framework"},
		{ID: "3", Content: "write a shop app"},
	}
	templates.WriteHello(w, items)
}

func (s *Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	go readLoop(conn, s.requests)
	writeLoop(conn, s.responses)
}

func writeLoop(conn *websocket.Conn, responses chan Message) {
	for {
		message := <-responses
		writer, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Sending %v", message)
		templates.WriteMessage(writer, message.ID, message.Content) //templates.Body(fmt.Sprint(i)))
		err = writer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Event struct {
	Type string `json:"event"`
	ID   string `json:"id"`
}

func readLoop(conn *websocket.Conn, requests chan Event) {
	for {
		_, reader, err := conn.NextReader()
		if err != nil {
			log.Fatal(err)
		}
		decoder := json.NewDecoder(reader)
		e := Event{}
		decoder.Decode(&e)
		requests <- e
	}
}
