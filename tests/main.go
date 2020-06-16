package main

import (
	"bufio"
	"github.com/discoriver/textbox"
	"os"
	"strings"
)

func main() {
	console := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout))
	player := &textbox.Player{Console: console}
	rootMap := &textbox.Map{Name: "map1"}

	newGameWorld(rootMap)
	start, err := rootMap.FindRoom([]byte{1, 1, 1})
	if err != nil {
		player.Println("Room failure.")
		os.Exit(1)
	}

	player.CurrentRoom = start

	Run(player)
}

func Run(p *textbox.Player) {
	p.Println("Welcome to this text adventure.")

	p.CurrentRoom.Enter()
	p.Look(true)

	for {
		p.Printf(">")

		cmd, err := p.Console.ReadString('\n')
		if err != nil {
			p.Println("error reading console")
			break
		}
		cmd = strings.ToUpper(strings.Trim(cmd, "\n"))

		cmd = p.VerbAliasReplace(cmd)

		if cmd == "QUIT" {
			p.Println("Thanks for playing!")
			break
		}

		p.ExecuteCommand(cmd)
	}
}
