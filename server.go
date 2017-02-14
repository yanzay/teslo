package teslo

import (
	"io"
	"net/http"

	"github.com/yanzay/log"
)

type Message struct {
	ID      string
	Content string
}

type Responder interface {
	Respond(string, string)
}

type Server struct {
	InitSession  func(string)
	CloseSession func(string)
	Render       func(io.Writer)
	Addr         string
	handlers     map[string]func(*Session, *Event)
}

func NewServer() *Server {
	return &Server{
		Addr:     ":8080",
		handlers: make(map[string]func(*Session, *Event)),
	}
}

func (s *Server) Start() {
	http.HandleFunc("/", s.IndexHandler)
	http.HandleFunc("/ws", s.WSHandler)
	log.Infof("Starting server at %s", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func (s *Server) Subscribe(id string, handler func(*Session, *Event)) {
	s.handlers[id] = handler
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w)
}

func (s *Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Can't update to WebSocket protocol: %s", err)
		return
	}
	session := NewSession(s, conn)
	s.InitSession(session.ID)
	session.Start()
}
