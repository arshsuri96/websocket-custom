package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Clients ClientList
	sync.RWMutex
}

// factory functions
func Newmanager() *Manager {
	return &Manager{
		Clients: make(ClientList),
	}
}

func (m *Manager) serverWs(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Close()

}
