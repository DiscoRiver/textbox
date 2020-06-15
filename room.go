package textbox

type Room struct {
	Position []byte
	Name string
	Description string
	EnterFunc func(*Player)
	ExitFunc func(direction string) bool
	ExitBlocked []string
	ObjectContainer
}



