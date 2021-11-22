package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) spaceSnakeEvent() {
	graphics.SpaceSnakeEvent()
	switch user.Input() {
	case "1":
		fmt.Println("test the mission specialist")
	case "2":
		fmt.Println("Lose some of the passengers")
	default:
		fmt.Println("lose a crew member")
	}
}
