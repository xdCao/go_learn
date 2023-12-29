package ipc

import (
	"encoding/json"
	"sync"
)

type CenterServer struct {
	servers map[string]Server
	players []*Player
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) AddPlayer(params string) error {
	player := &Player{}

	err := json.Unmarshal([]byte(params), player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) RemovePlayer(name string) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for idx, v := range server.players {
		if v.Name == name {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
				return
			}
			if idx == 0 {
				server.players = server.players[1:]
			} else if idx == len(server.players)-1 {
				server.players = server.players[:len(server.servers)-2]
			} else {
				server.players = append(server.players[:idx-1], server.players[idx+1:]...)
			}
		}
	}
}

func (server *CenterServer) ListPlayers() string {
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	b, _ := json.Marshal(server.players)
	return string(b)
}

func (server *CenterServer) broadcast(params string) {
	message := &Message{}
	err := json.Unmarshal([]byte(params), message)
	if err != nil {
		return
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for _, v := range server.players {
		v.mq <- message
	}

}

func (server *CenterServer) Handle(method, params string) *Response {
	switch method {
	case "addplayer":
		err := server.AddPlayer(params)
		if err != nil {
			return &Response{Code: err.Error()}
		}
	case "removeplayer":
		server.RemovePlayer(params)
		return &Response{Code: "200"}
	case "listplayer":
		players := server.ListPlayers()
		return &Response{"200", players}
	case "broadcast":
		server.broadcast(params)
		return &Response{Code: "200"}
	default:
		return &Response{Code: "404", Body: method + ":" + params}
	}
	return &Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
