package main

import (
	"fmt"
	"log"
	"net/http"

	deliveryHttp "food/internal/delivery/http"
	"food/internal/usecase"
)

func main() {
	// Initialize the PingUseCase.
	pingUC := usecase.NewPingUseCase()
	// Create a new PingHandler with the PingUseCase.
	pingHandler := deliveryHttp.NewPingHandler(pingUC)

	// Set up the HTTP route for the "/ping" endpoint.
	http.HandleFunc("/ping", pingHandler.Ping)

	// Define the port for the server to listen on.
	port := ":8081"
	fmt.Println("Server is running on port", port)
	// Start the HTTP server and log any errors.
	log.Fatal(http.ListenAndServe(port, nil))
}
