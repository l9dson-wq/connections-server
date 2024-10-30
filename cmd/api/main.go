package main

import (
    "log"
    "gameServer/internal/server"
)

func main() {
    log.Println("Server started...")

    //calling the server to start it
    server.StartServer()
}
