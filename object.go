package textbox

import (
	"errors"
	"fmt"
	"strings"
)

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

// Return true if the string matches the object.
// TODO: I don't like this logic. Placing it here for now.
func (o *Object) RespondTo(args []string) bool {
	str := strings.Join(args, " ")
	// match all adjectives
	for _, adj := range o.Adjectives {
		if i := strings.Index(str, strings.ToUpper(adj)+" "); i == 0 {
			str = str[len(adj)+1:]
		}
	}
	if str == strings.ToUpper(o.Name) {
		return true
	}
	for _, alias := range o.Aliases {
		if str == strings.ToUpper(alias) {
			return true
		}
	}
	return false
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

func (c *ObjectContainer) AddObject(objs ...*Object) {
	c.Objects = append(c.Objects, objs...)
}

func (c *ObjectContainer) RemoveObject(obj *Object) {
	loc := -1
	for i, val := range c.Objects {
		if val == obj {
			loc = i
			break
		}
	}
	c.Objects = append(c.Objects[:loc], c.Objects[loc+1:]...)
}

func (c *ObjectContainer) CollectObjectNames() (res string, err error) {
	count := 0
	res = ""
	for i, obj := range c.Objects {
		if obj.Fixture {
			continue
		}
		res += obj.GetName()
		count++
		if i < len(c.Objects)-2 {
			res += ", "
		} else if i == len(c.Objects)-2 {
			res += ", and "
		}
	}
	if count == 0 {
		err = errors.New("no objects found")
	}
	return
}

// TODO: This is the only function that relies on RespondTo. Find a more intuitive way to implement object detection.
func (c *ObjectContainer) FindObject(args []string) *Object {
	for _, obj := range c.Objects {
		if obj.RespondTo(args) {
			return obj
		}
	}
	return nil
}

