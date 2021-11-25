package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) foodSpoilageEvent() {
	graphics.FoodSpoilageEvent()
	switch user.Input() {
	case "1":
		if gs.Engineer.TestSkill(50) {
			fmt.Println("Your engineer managed to repair the malfunctioning systems.")
			return
		}
		fmt.Println("The malfunction systems were not repaired.")
	case "2":
		if gs.RepairParts >= 5 {
			fmt.Println("You use your mission's repair parts.")
			gs.RepairParts -= 5
			return
		} else {
			fmt.Println("You do not have enough repair parts!")
		}
	default:
		fmt.Println("Not all of the food will fit, some of it has to be jettisoned after spoiling.")
	}

	fmt.Println("You lose mission food supplies.")
	if gs.Food < 5 {
		gs.Food = 0
		return
	}
	gs.Food -= 5
}
