package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type Server struct {
	mu    sync.Mutex
	conns map[*websocket.Conn]bool
}

type IncomingMessage struct {
	Message string                 `json:"message"`
	Headers map[string]interface{} `json:"HEADERS"`
}

func main() {
	r := mux.NewRouter()
	server := newServer()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", homeRoute)

	r.Handle("/ws", websocket.Handler(server.handleWs))

	http.ListenAndServe(":8080", r)
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWs(ws *websocket.Conn) {
	log.Println("Established connection with ", ws.RemoteAddr())
	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.conns, ws)
		s.mu.Unlock()
		log.Println("Connection closed for ", ws.RemoteAddr())
	}()

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error reading buffer:", err)
			continue
		}
		msg := buf[:n]

		var incomingMessage IncomingMessage
		if err := json.Unmarshal(msg, &incomingMessage); err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		htmlMessage := formatAsHTML(incomingMessage)
		s.broadcast([]byte(htmlMessage))
	}
}

func formatAsHTML(msg IncomingMessage) string {
	hxTarget, ok := msg.Headers["HX-Target"].(string)
	if !ok {
		return "Invalid HX-Target format"
	}

	return `<div id="` + hxTarget + `" hx-swap-oob="beforeend"><br>` + msg.Message + `</div>`
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println("write error ", err)
			}
		}(ws)
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}
