package ipc

import "fmt"

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	Room  int    "room"

	mq chan *Message
}

type Message struct {
	From string "from"
	To   string "to"
	Body string "body"
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-m
			fmt.Printf("msg: %v\n", msg)
		}
	}(player)
	return player
}
