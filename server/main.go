package main

import (
	"log"
	"net/http"
	"server/handlers" // Adjust this import path based on your project's structure.
)

func main() {
	go handlers.HandleMessages() // Start listening for incoming chat messages

	// Register your WebSocket endpoint (assuming you named it HandleWebSocketConnection)
	http.HandleFunc("/ws", handlers.HandleWebSocketConnection)

	// Start your server (choose a port that suits you, 8080 for this example)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
