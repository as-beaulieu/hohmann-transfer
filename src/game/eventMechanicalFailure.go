package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) mechanicalFailureEvent() {
	graphics.MechanicalFailureEvent()
	switch user.Input() {
	case "1":
		fmt.Println("Test the crew")
		if gs.Engineer.TestSkill(50) {
			fmt.Println("Your engineer managed to repair the malfunctioning systems.")
			return
		}
		fmt.Println("The malfunctioning systems were not repaired.")
	case "2":
		if gs.RepairParts >= 5 {
			fmt.Println("You use your mission's repair parts.")
			gs.LoseRepairParts(5)
			return
		} else {
			fmt.Println("You do not have enough repair parts!")
		}
	default:
		fmt.Println("Do nothing.")
	}

	fmt.Println("The mechanical issues strain other systems on the ship,")
	fmt.Println("Causing other systems to fail.")
	gs.LosePassengers(Random(10))
	gs.LoseFood(Random(5))
	gs.LoseOxygen(Random(10))
	gs.LoseWater(Random(10))
}
