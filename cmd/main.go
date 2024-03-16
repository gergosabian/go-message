package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

func main() {
	server := NewServer()
	chatroom := NewChatroom()
	page := NewPage(chatroom)

	http.Handle("/chat", websocket.Handler(func(ws *websocket.Conn) {
		server.HandleChat(ws, chatroom)
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("index").ParseFiles("templates/index.html", "templates/chatroom.html"))
		if err := t.Execute(w, *page); err != nil {
			fmt.Println("Error executing template:", err)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}

type Server struct {
	conns map[*websocket.Conn]bool
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleChat(ws *websocket.Conn, chatroom *Chatroom) {
	fmt.Println("New connection:", ws.RemoteAddr())
	s.mutex.Lock()
	s.conns[ws] = true
	s.mutex.Unlock()
	defer func() {
		s.mutex.Lock()
		delete(s.conns, ws)
		s.mutex.Unlock()
	}()

	for {
		var message Message

		if err := websocket.JSON.Receive(ws, &message); err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed:", ws.RemoteAddr())
				break
			}
			fmt.Println("Error receiving message:", err)
			continue
		}

		chatroom.Messages = append(chatroom.Messages, message)

		html := bytes.Buffer{}
		t := template.Must(template.New("chatroom").ParseFiles("templates/chatroom.html"))

		if err := t.Execute(&html, chatroom); err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		for conn := range s.conns {
			if err := websocket.Message.Send(conn, html.String()); err != nil {
				fmt.Println("Error sending message:", err)
				continue
			}
		}
	}
}

type Page struct {
	Chatroom *Chatroom
}

func NewPage(chatroom *Chatroom) *Page {
	return &Page{
		Chatroom: chatroom,
	}
}

type Chatroom struct {
	Messages []Message
}

func NewChatroom() *Chatroom {
	return &Chatroom{
		Messages: make([]Message, 0),
	}
}

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

func NewMessage(sender, content string) *Message {
	return &Message{
		Sender:  sender,
		Content: content,
	}
}
