package kbmap

import (
	"errors"
	"github.com/driusan/de/demodel"
	"golang.org/x/mobile/event/key"
)

// A Map maps a keystroke to a command. It performs a command, and then
// returns a new map which represents the keyboard mapping to be used
// for the next keystroke.
type Map interface {
	HandleKey(key.Event, *demodel.CharBuffer) (Map, error)
}

var Invalid error = errors.New("Invalid keyboard map.")
var ExitProgram error = errors.New("Keystroke wants to exit the program.")
var ScrollDown error = errors.New("Keystroke wants to scroll the window down.")
var ScrollUp error = errors.New("Keystroke wants to scroll the window up.")
var ScrollLeft error = errors.New("Keystroke wants to scroll the window left.")
var ScrollRight error = errors.New("Keystroke wants to scroll the window right.")

type defaultMaps uint

const (
	NormalMode = defaultMaps(iota)
	InsertMode
	DeleteMode
	TagMode
)

func (m defaultMaps) HandleKey(e key.Event, buff *demodel.CharBuffer) (Map, error) {
	switch m {
	case NormalMode:
		return normalMap(e, buff)
	case InsertMode:
		return insertMap(e, buff)
	case DeleteMode:
		return deleteMap(e, buff)
	case TagMode:
		return tagMap(e, buff)
	}
	return nil, Invalid

}