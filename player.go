package textbox

import "bufio"

type Player struct {
	Console *bufio.ReadWriter
	CurrentMap *Map
	CurrentRoom *Room
	ObjectContainer
}

func (p *Player) MoveRoom(direction string) (*Room, error) {
	for i := range p.CurrentRoom.ExitBlocked {
		if direction == p.CurrentRoom.ExitBlocked[i] {
			return nil, ErrBlocked
		}
	}

	position := p.CurrentRoom.Position
	switch direction {
	case "NORTH":
		position[1]++
	case "SOUTH":
		position[1]--
	case "EAST":
		position[0]++
	case "WEST":
		position[0]--
	case "UP":
		position[2]--
	case "DOWN":
		position[2]++
	default:
		return nil, ErrCannotCompute
	}

	newRoom, err := p.CurrentMap.FindRoom(position)
	if err == ErrBlocked {
		return nil, err
	}

	return newRoom, nil
}


