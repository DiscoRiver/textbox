package main

import "github.com/discoriver/textbox"

func newGameWorld(m *textbox.Map) {
	Room11 := &textbox.Room{
		Position: []byte{1, 1, 1},
		Name: "Lobby",
		Description: "You find yourself in a hotel lobby.",
		ExitBlocked: []string{"SOUTH"},
	}

	trapdoor := &textbox.Object{
		Name:      "Trapdoor",
		Carryable: false,
		Openable:  true,
		Aliases:   []string{"door"},
	}

	Room11.AddObject(trapdoor)
	m.AddRoom(Room11)
}
