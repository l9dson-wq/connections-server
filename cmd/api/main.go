package main

import (
    "log"
    "gameServer/internal/server"
)

func main() {
    log.Println("Server started...")

    // difining the list of IDs
    playerIDs := []string{"1", "2", "3", "4", "5"}

    // difining the URL of the server
    URL := "http://localhost:8090/connect"

    // calling the connections function to simulate them
    go server.SimulateConnections(playerIDs, URL)

    //calling the server to start it
    server.StartServer()
}
