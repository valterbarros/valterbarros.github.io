package editor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"valterbarros/blog-tools/internal/markdown"
	"valterbarros/blog-tools/internal/utils"

	"github.com/gorilla/websocket"
)

// DTO
type Action struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (a *Action) toGo(msg []byte) {
	if err := json.Unmarshal(msg, a); err != nil {
		log.Println("Erro ao ler JSON:", err)
	}
}

func NewAction() *Action {
	return &Action{}
}

// DTO

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao fazer upgrade: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Cliente conectado via WebSocket!")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem (cliente desconectou?): %v\n", err)
			break
		}

		// fmt.Printf("Mensagem recebida: %s\n", string(msg))

		ac := NewAction()
		ac.toGo(msg)

		switch ac.Type {
		case "change":
			result := markdown.MarkToHTML([]byte(ac.Value))
			// send markdown in html
			err = conn.WriteMessage(messageType, []byte(result))
		case "update":
			path, _ := utils.GetFromRoot("articles/_posts/2026-3-13-mudancas.md")
			utils.SaveFile(path, ac.Value)
		}

		fmt.Printf("action %#v", ac.Type)

		if err != nil {
			log.Printf("Erro ao escrever mensagem: %v\n", err)
			break
		}
	}
}
