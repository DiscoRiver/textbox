/*
The textbox package is a set of tools for creating complex text adventures.
*/
package textbox

import "errors"

var (
	ErrBlocked = errors.New("You cannot exit this way.")
	ErrCannotCompute = errors.New("I don't understand that instruction.")
)
