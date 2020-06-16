package textbox

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Player struct {
	Console *bufio.ReadWriter
	CurrentMap *Map
	CurrentRoom *Room
	ObjectContainer
}

func (p *Player) requestMoveRoom(direction string) (*Room, error) {
	for i := range p.CurrentRoom.ExitBlocked {
		if direction == p.CurrentRoom.ExitBlocked[i] {
			return nil, errRouteBlocked
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
		return nil, errCannotCompute
	}

	newRoom, err := p.CurrentMap.FindRoom(position)
	if err == errRouteBlocked {
		return nil, err
	}

	return newRoom, nil
}

func (p *Player) Go(args []string) bool {
	if newRoom, err := p.requestMoveRoom(args[0]); err != nil && newRoom != nil {
		p.CurrentRoom.Leave()
		p.CurrentRoom = newRoom
		p.Look(!newRoom.visited)
		newRoom.Enter()
	} else {
		p.Println(errRouteBlocked.Error())
	}
	return true
}

func (p *Player) Look(printDescription bool) bool {
	p.Println(p.CurrentRoom.Name)
	if printDescription && p.CurrentRoom.Description != "" {
		p.Println(p.CurrentRoom.Description)
	}
	if objectString, err := p.CurrentRoom.CollectObjectNames(); err == nil {
		p.Println("There is " + objectString + " here.")
	}
	return true
}

func (p *Player) LookAt(args []string) bool {
	if len(args) == 0 {
		return p.Look(true)
	}
	if object := p.FindNearObject(args); object != nil {
		p.Println(object.GetDescription())
	} else {
		p.Printf("I don't see any %v here.\n", strings.Join(args, " "))
	}
	return true
}

func (p *Player) Take(args []string) bool {
	if object := p.CurrentRoom.FindObject(args); object != nil {
		if object.Carryable && !object.Fixture {
			p.CurrentRoom.RemoveObject(object)
			p.AddObject(object)
			p.Println("Taken.")
		} else {
			p.Println("This can't be taken.")
		}
	} else {
		p.Printf("I don't see any %v here.\n", strings.Join(args, " "))
	}
	return true
}

func (p *Player) Push(object *Object) bool {
	if object != nil {
		if cb := object.Verbs["PUSH"]; cb != nil {
			cb(object, p)
		} else {
			p.Println("You can't push this.")
		}
	} else {
		p.Println("I don't know what you're referring to.")
	}
	return true
}

func (p *Player) Pull(object *Object) bool {
	if object != nil {
		if cb := object.Verbs["PULL"]; cb != nil {
			cb(object, p)
		} else {
			p.Println("You can't pull this.")
		}
	} else {
		p.Println("I don't know what you're referring to.")
	}
	return true
}

func (p *Player) LookUnder(object *Object) bool {
	if object != nil {
		if cb := object.Verbs["LOOK UNDER"]; cb != nil {
			cb(object, p)
		} else {
			p.Println("You don't see anything out of the ordinary.")
		}
	} else {
		p.Println("I don't know what you are referring to.")
	}
	return true
}

func (p *Player) Drop(args []string) bool {
	if object := p.FindObject(args); object != nil {
		p.RemoveObject(object)
		p.CurrentRoom.AddObject(object)
		p.Println("Dropped.")
	} else {
		p.Printf("I don't see any %v here.\n", strings.Join(args, " "))
	}
	return true
}

func (p *Player) Open(args []string) bool {
	if obj := p.FindNearObject(args); obj != nil {
		if obj.Openable {
			if !obj.Open {
				obj.Open = true
				p.Println("Opened.")
			} else {
				p.Println("Already open.")
			}
		} else {
			p.Println("I can't open that.")
		}
	} else {
		p.Printf("I don't see any %v here.\n", strings.Join(args, " "))
	}
	return true
}

func (p *Player) Close(args []string) bool {
	if obj := p.FindNearObject(args); obj != nil {
		if obj.Openable {
			if obj.Open {
				obj.Open = false
				p.Println("Closed.")
			} else {
				p.Println("Already closed.")
			}
		} else {
			p.Println("I can't close that.")
		}
	} else {
		p.Printf("I don't see any %v here.\n", strings.Join(args, " "))
	}
	return true
}

func (p *Player) Wait() bool {
	p.Println("Time passes.")
	return true
}

func (p *Player) Help(args []string) bool {
	p.Println("This is a text adventure game.")
	p.Println("The game only understands very simple single-verb, single-object sentences, for instance: PICK UP HAT, or OPEN DOOR etc.")
	p.Println("The Verbs this game understands are: LOOK, LOOK AT, LOOK UNDER, PUSH, PULL, TAKE, DROP, WAIT, OPEN, CLOSE and INVENTORY.")
	p.Println("Directions are: NORTH, SOUTH, EAST, WEST, UP, DOWN, IN and OUT.")
	p.Println("There are also many aliases for verbs and directions.")
	return true
}

func (p *Player) ExecuteCommand(command string) bool {
	verbMap := map[string]func([]string) bool{
		"GO":         func(args []string) bool { return p.Go(args) },
		"LOOK":       func(args []string) bool { return p.Look(true) },
		"LOOK AT":    func(args []string) bool { return p.LookAt(args) },
		"TAKE":       func(args []string) bool { return p.Take(args) },
		"PUSH":       func(args []string) bool { return p.Push(p.FindNearObject(args)) },
		"PULL":       func(args []string) bool { return p.Pull(p.FindNearObject(args)) },
		"LOOK UNDER": func(args []string) bool { return p.LookUnder(p.CurrentRoom.FindObject(args)) },
		"DROP":       func(args []string) bool { return p.Drop(args) },
		"OPEN":       func(args []string) bool { return p.Open(args) },
		"WAIT":       func(args []string) bool { return p.Wait() },
		"CLOSE":      func(args []string) bool { return p.Close(args) },
		"INVENTORY":  func(args []string) bool { return p.Inventory(args) },
		"XYZZY":      func(args []string) bool { return true },
		"HELP":       func(args []string) bool { return p.Help(args) },
	}
	delegated := false
	// we need to make sure to sort the verbs by length first:
	verbs := []string{} // make([]string, len(verbMap))
	for verb := range verbMap {
		verbs = append(verbs, verb)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(verbs)))
	for _, verb := range verbs {
		fn := verbMap[verb]
		if verb == command {
			delegated = fn(make([]string, 0))
		} else if i := strings.Index(command, verb+" "); i == 0 {
			command = command[len(verb)+1:]
			delegated = fn(strings.Split(command, " "))
		}
	}
	if !delegated {
		p.Println("Sorry, what?")
	}
	return delegated
}

var verbAliasMap = map[string][]string{
	"GO NORTH":   {"N", "NORTH"},
	"GO SOUTH":   {"S", "SOUTH"},
	"GO WEST":    {"W", "WEST"},
	"GO EAST":    {"E", "EAST"},
	"GO IN":      {"IN", "INSIDE", "ENTER"},
	"GO OUT":     {"OUT", "OUTSIDE", "LEAVE"},
	"GO UP":      {"UP"},
	"GO DOWN":    {"DOWN"},
	"LOOK AT":    {"EXAMINE", "INSPECT", "X"},
	"LOOK UNDER": {"LOOK BENEATH", "LOOK BELOW"},
	"TAKE":       {"PICK UP", "GET"},
	"DROP":       {"THROW"},
	"INVENTORY":  {"I"},
	"WAIT":       {"Z"},
}

func (p *Player) Inventory(args []string) bool {
	if objectString, err := p.CollectObjectNames(); err == nil {
		p.Println("You are carrying " + objectString + ".")
	} else {
		p.Println("You are empty handed.")
	}
	return true
}

func (p *Player) FindNearObject(args []string) *Object {
	if object := p.CurrentRoom.FindObject(args); object != nil {
		return object
	}

	if object := p.FindObject(args); object != nil {
		return object
	}
	return nil
}

func (p *Player) Println(line string) {
	p.Printf(line + "\n")
}

func (p *Player) Printf(format string, args ...interface{}) {
	line := fmt.Sprintf(format, args...)
	p.Console.WriteString(line)
	p.Console.Flush()
}

var VerbAliasMap = map[string][]string{
	"GO NORTH":   {"N", "NORTH"},
	"GO SOUTH":   {"S", "SOUTH"},
	"GO WEST":    {"W", "WEST"},
	"GO EAST":    {"E", "EAST"},
	"GO IN":      {"IN", "INSIDE", "ENTER"},
	"GO OUT":     {"OUT", "OUTSIDE", "LEAVE"},
	"GO UP":      {"UP"},
	"GO DOWN":    {"DOWN"},
	"LOOK AT":    {"EXAMINE", "INSPECT", "X"},
	"LOOK UNDER": {"LOOK BENEATH", "LOOK BELOW"},
	"TAKE":       {"PICK UP", "GET"},
	"DROP":       {"THROW"},
	"INVENTORY":  {"I"},
	"WAIT":       {"Z"},
}

func (p *Player) VerbAliasReplace(cmd string) string {
	for verb, aliases := range VerbAliasMap {
		for _, alias := range aliases {
			if alias == cmd {
				return verb
			} else if i := strings.Index(cmd, verb+" "); i == 0 {
				return cmd
			} else if i := strings.Index(cmd, alias+" "); i == 0 {
				return strings.Replace(cmd, alias+" ", verb+" ", 1)
			}
		}
	}
	return cmd
}



