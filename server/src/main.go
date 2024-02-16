package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kaaaaazuya/chatapp/src/domain"
	"github.com/Kaaaaazuya/chatapp/src/handlers"
)

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}