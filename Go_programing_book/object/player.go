package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	player "xdCao/golearn/goprogramming/player"
)

var manager *player.Manager
var id int = 1

func main() {

	fmt.Println(`Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)

	manager = player.NewManager(make([]player.Music, 0))

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter commmand...")
		line, _, _ := r.ReadLine()
		str_line := string(line)
		if str_line == "q" || str_line == "e" {
			break
		}
		tokens := strings.Split(str_line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("not support command")
		}
	}
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < manager.Len(); i++ {
			v, _ := manager.Get(i)
			fmt.Println(i, ":", v)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			manager.Add(&player.Music{
				Id:       strconv.Itoa(id),
				Name:     tokens[2],
				Artist:   tokens[3],
				Location: tokens[4],
				FileType: tokens[5],
			})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			manager.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	music := manager.Find(tokens[1])
	if music == nil {
		fmt.Println("music not found")
		return
	}
	player.Play(music.Location, music.FileType)
}
