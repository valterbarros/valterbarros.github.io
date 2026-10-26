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
	File  string `json:"file"`
}

func parseAction(msg []byte) (*Action, error) {
	var a Action

	if err := json.Unmarshal(msg, &a); err != nil {
		return nil, err
	}

	return &a, nil
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

	fmt.Println("Connected via websocket!")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro to read message: %v\n", err)
			break
		}

		ac, err := parseAction(msg)

		switch ac.Type {
		case "change":
			result := markdown.MarkToHTML([]byte(ac.Value))
			// send markdown in html
			// TODO: remove YAML data "---"
			err = conn.WriteMessage(messageType, []byte(result))
		case "update":
			// TODO: it is fixed filename
			path, _ := utils.GetFromRoot("articles/_posts/" + ac.File)
			utils.SaveFile(path, ac.Value)
		}

		if err != nil {
			log.Printf("Error to write response: %v\n", err)
			break
		}
	}
}
