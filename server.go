package teslo

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

type Responder interface {
	Respond(string, string)
}

type Server struct {
	Render    func(http.ResponseWriter)
	Addr      string
	requests  chan Event
	responses chan Message
	handlers  map[string]func(Responder, *Event)
}

func NewServer() *Server {
	return &Server{
		Addr:      ":8080",
		requests:  make(chan Event, 0),
		responses: make(chan Message, 0),
		handlers:  make(map[string]func(Responder, *Event)),
	}
}

func (s *Server) Start() {
	go s.eventHandler()
	http.HandleFunc("/", s.IndexHandler)
	http.HandleFunc("/ws", s.WSHandler)
	log.Infof("Starting server at %s", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func (s *Server) eventHandler() {
	for {
		request := <-s.requests
		log.Println(request)
		if s.handlers[request.ID] != nil {
			s.handlers[request.ID](s, &request)
		} else {
			for _, parent := range request.Parents {
				if s.handlers[parent] != nil {
					s.handlers[parent](s, &request)
					break
				}
			}
		}
	}
}

func (s *Server) Respond(id string, content string) {
	s.responses <- Message{ID: id, Content: content}
}

func (s *Server) Subscribe(id string, handler func(Responder, *Event)) {
	s.handlers[id] = handler
}

var upgrader = websocket.Upgrader{}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w)
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
		templates.WriteMessage(writer, message.ID, message.Content)
		err = writer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Event struct {
	Type    string   `json:"event"`
	ID      string   `json:"id"`
	Parents []string `json:"parents"`
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
