package textbox

type Room struct {
	Position []byte
	Name string
	Description string
	visited bool
	EnterFunc func(*Player)
	ExitFunc func(direction string) bool
	ExitBlocked []string // NORTH, EAST, SOUTH, WEST
	ObjectContainer
}

func (r *Room) Enter() {
	r.visited = true
}


func (r *Room) Leave() {
	// Do we need exit logic?
}

