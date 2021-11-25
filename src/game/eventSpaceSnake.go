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
		if gs.Specialist.TestSkill(50) {
			fmt.Println("Your mission specialist capture the snake and added it to the lab.")
			return
		}
	case "2":
		fmt.Println("The snake was found after three days of searching the ship.")
		fmt.Println("The snake had several lumps in it. Your crew is certain that they may be the connected ")
		fmt.Println("to the missing passengers reported.")
		gs.Passengers -= Random(5) + 1
	default:
		member, err := gs.WhoToKill()
		if err != nil {
			gs.MissionFailure = true
			return
		}
		fmt.Printf("You lose your Mission %v Officer has died of space snakebite", member)
		gs.KillCrewMember(member)
	}
}
