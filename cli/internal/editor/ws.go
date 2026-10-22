package editor

import (
	"fmt"
	"log"
	"net/http"
	"valterbarros/blog-tools/internal/markdown"

	"github.com/gorilla/websocket"
)

// Upgrader converte a requisição HTTP normal em um WebSocket persistente
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permite conexões de qualquer origem (ideal para dev local)
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. Faz o upgrade da conexão HTTP para WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao fazer upgrade: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Cliente conectado via WebSocket!")

	// 2. Loop para escutar mensagens do cliente indefinidamente
	for {
		// Lê a mensagem enviada pelo cliente
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem (cliente desconectou?): %v\n", err)
			break
		}

		fmt.Printf("Mensagem recebida: %s\n", string(msg))

		result := markdown.MarkToHTML(msg)

		// send markdown in html
		err = conn.WriteMessage(messageType, []byte(result))
		if err != nil {
			log.Printf("Erro ao escrever mensagem: %v\n", err)
			break
		}
	}
}
