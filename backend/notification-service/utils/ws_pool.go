package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	wsClients      = make(map[int]*websocket.Conn)
	wsClientsMutex = sync.RWMutex{}
)

func AddWebSocketConn(userID int, conn *websocket.Conn) {
	wsClientsMutex.Lock()
	defer wsClientsMutex.Unlock()
	wsClients[userID] = conn
}

func GetWebSocketConn(userID int) *websocket.Conn {
	wsClientsMutex.RLock()
	defer wsClientsMutex.RUnlock()
	return wsClients[userID]
}

func RemoveWebSocketConn(userID int) {
	wsClientsMutex.Lock()
	defer wsClientsMutex.Unlock()
	delete(wsClients, userID)
}
