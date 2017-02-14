package teslo

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/yanzay/log"
	"github.com/yanzay/teslo/templates"
)

var upgrader = websocket.Upgrader{}

type Session struct {
	ID        string
	conn      *websocket.Conn
	requests  chan Event
	responses chan Message
	server    *Server
	close     chan struct{}
}

func NewSession(server *Server, conn *websocket.Conn) *Session {
	return &Session{
		ID:        uuid.NewV4().String(),
		server:    server,
		conn:      conn,
		requests:  make(chan Event, 0),
		responses: make(chan Message, 0),
		close:     make(chan struct{}, 0),
	}
}

func (s *Session) Start() {
	log.Debugf("Starting session: %s", s.ID)
	defer s.server.CloseSession(s.ID)
	go s.writeLoop()
	go s.readLoop()
	s.eventHandler()
}

func (s *Session) Close() {
	log.Debugf("Closing session: %s", s.ID)
	close(s.close)
}

func (s *Session) Respond(id string, content string) {
	s.responses <- Message{ID: id, Content: content}
}

func (s *Session) eventHandler() {
	for {
		select {
		case <-s.close:
			log.Debugf("Exiting eventHandler for session: %s", s.ID)
			return
		case request := <-s.requests:
			log.Debugf("Sending request: %v", request)
			if s.server.handlers[request.ID] != nil {
				s.server.handlers[request.ID](s, &request)
			} else {
				for _, parent := range request.Parents {
					if s.server.handlers[parent] != nil {
						s.server.handlers[parent](s, &request)
						break
					}
				}
			}
		}
	}
}

func (s *Session) writeLoop() {
	for {
		select {
		case <-s.close:
			log.Debugf("Exiting writeLoop for session: %s", s.ID)
			return
		case message := <-s.responses:
			writer, err := s.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Errorf("Can't get next writer for websocket: %s", err)
				s.Close()
				return
			}
			log.Debugf("Sending response: %v", message)
			templates.WriteMessage(writer, message.ID, message.Content)
			err = writer.Close()
			if err != nil {
				log.Errorf("Can't write to websocket: %s", err)
				s.Close()
				return
			}
		}
	}
}

type Event struct {
	Type    string   `json:"event"`
	ID      string   `json:"id"`
	Parents []string `json:"parents"`
}

func (s *Session) readLoop() {
	for {
		select {
		case <-s.close:
			log.Debugf("Exiting readLoop for session: %s", s.ID)
			return
		default:
			_, reader, err := s.conn.NextReader()
			if err != nil {
				log.Errorf("Can't get websocket reader: %s", err)
				s.Close()
				return
			}
			decoder := json.NewDecoder(reader)
			e := Event{}
			decoder.Decode(&e)
			s.requests <- e
		}
	}
}
