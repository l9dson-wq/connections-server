package server

import (
	"fmt"
	"net/http"
	"sync"
    _ "log"
)

type Server struct {
    players map[string]bool
    mutex   sync.Mutex
}

func (s *Server) Broadcast(message string) {
    fmt.Printf(message)
}

func (s *Server) ConnectPlayer (w http.ResponseWriter, r *http.Request) {
    playerID := r.URL.Query().Get("id")
    if playerID == "" {
        http.Error(w, "The 'id' parameter is missing", http.StatusBadRequest)
        return
    }

    // lock the thread
    s.mutex.Lock()
    // unlock the thread
    defer s.mutex.Unlock()

    // if an ID was send trough the URL we connect the player to the session.
    if !s.players[playerID]{
        s.players[playerID] = true
        s.Broadcast(fmt.Sprintf("Player %s has joined the session. \n", playerID))
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Player %s connected\n", playerID)))
}

func (s *Server) DisconnectPlayer (w http.ResponseWriter, r *http.Request) {
    playerID := r.URL.Query().Get("id")
    if playerID == "" {
        http.Error(w, "The 'id' parameter is missing", http.StatusBadRequest)
        return
    }

    s.mutex.Lock()
    defer s.mutex.Unlock()

    if s.players[playerID]{
        delete(s.players, playerID)
        s.Broadcast(fmt.Sprintf("Player %s has leave the session.\n", playerID))
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Player %s disconnected", playerID)))
}
