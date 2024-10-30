package server

import (
	"log"
	"net/http"
)

func StartServer() {
    server := &Server{players: make(map[string]bool)}

    http.HandleFunc("/connect", server.ConnectPlayer)
    http.HandleFunc("/disconnect", server.DisconnectPlayer)

    log.Println("Server started in 8090 port")
    http.ListenAndServe(":8090", nil)
}
