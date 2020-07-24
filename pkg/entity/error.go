package entity

import "math/rand"

//Error define if response will be error
type Error struct {
	Chance int

	Definition Definition
}

//Happened check if error happened
func (e Error) Happened() bool {
	return rand.Intn(100) <= e.Chance
}
