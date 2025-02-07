package main

import (
	"fmt"
	"log"
	"net/http"

	pingHttp "food/internal/delivery/http"
	"food/internal/usecase"
)

func main() {
	pingUC := usecase.NewPingUseCase()
	pingHandler := pingHttp.NewPingHandler(pingUC)

	http.HandleFunc("/ping", pingHandler.Ping)

	port := ":8081"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
