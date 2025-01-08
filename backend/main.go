package main

import (
	"log"
	"net/http"

	"github.com/natnael-alemayehu/chat-go/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handlers.HandleConnections)

	log.Print("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve error: ", err)
	}
}
