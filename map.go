package textbox

import (
	"bytes"
	"fmt"
)

/*
The Map struct is the top level container for the world. The hierarchy is as follows;

- A Map is a geographical "play area"
- A Map contains rooms.
- A Room contains objects & object containers.
- Objects can themselves, be object containers (e.g. a chest)
 */
type Map struct {
	Name string
	Rooms []*Room
}

// FindRoom determines if a room exists in the direction specified, and returns it, or nil on error.
func (m *Map) FindRoom(newPosition []byte) (*Room, error) {
	for i := range m.Rooms {
		if bytes.Equal(newPosition, m.Rooms[i].Position) {
			return m.Rooms[i], nil
		}
	}
	return nil, ErrBlocked
}

// AddRoom adds the specified room to the map. Errors if room position already exists.
func (m *Map) AddRoom(r *Room) error {
	for i := range m.Rooms {
		if bytes.Equal(r.Position, m.Rooms[i].Position) {
			return fmt.Errorf("could not add room to map, position [%v] already exists", r.Position)
		}
	}

	m.Rooms = append(m.Rooms, r)

	return nil
}
