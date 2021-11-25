package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) passengerMishapEvent() {
	graphics.PassengerMishapEvent()
	switch user.Input() {
	case "1":
		fmt.Println("Test the crew")
		if gs.Engineer.TestSkill(50) {
			fmt.Println("Your engineer managed to repair the malfunctioning systems.")
			return
		}
		fmt.Println("The malfunctioning passenger pod was not repaired.")
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

	fmt.Println("The passenger pod explosively decompresses, jettisoning all contents into space")
	gs.LosePassengers(Random(25))
	gs.LoseFood(Random(10))
	gs.LoseOxygen(Random(15))
	gs.LoseWater(Random(10))
}
