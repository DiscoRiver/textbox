/*
The textbox package is a set of tools for creating complex text adventures.
*/
package textbox

import "errors"

var (
	errRouteBlocked  = errors.New("You cannot go this way.")
	errCannotCompute = errors.New("I don't understand that instruction.")
)
