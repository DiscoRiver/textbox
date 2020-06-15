package textbox

import "fmt"

type Object struct {
	ID int // Reusable Object IDs
	Name string
	Description string
	Aliases []string
	Adjectives []string
	Verbs map[string]func(*Object, *Player)
	Openable bool
	Open bool
	Fixture bool
	Carryable bool
	ObjectContainer
}

func (o *Object) GetName() string {
	res := "a " + o.Name
	if o.Openable {
		if o.Open {
			res += " (open)"
		} else {
			res += " (closed)"
		}
	}
	return res
}

func (o *Object) GetDescription() string {
	res := ""
	if len(o.Description) > 0 {
		res = o.Description
	}
	if o.Openable {
		if o.Open {
			res += fmt.Sprintf("\nThe %v is open.", o.Name)
		} else {
			res += fmt.Sprintf("\nThe %v is closed.", o.Name)
		}
	}
	return res
}

type ObjectContainer struct {
	Objects []*Object
}

