package startPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func procurement(gs *game.State) {
	for {
		graphics.ClearScreen(gs.OperatingSystem)
		gs.StatusScreen()
		fmt.Println("|                                                               |")
		fmt.Println("  You should need a minimum of 30 units of fuel for the mission  ")
		fmt.Println("|                                                               |")
		fmt.Println("| (1) - Fuel (1 unit / Cost: 5)                                 |")
		fmt.Println("| (2) - Food (1 unit / Cost: 2)                                 |")
		fmt.Println("| (3) - Water (1 unit / Cost: 1)                                |")
		fmt.Println("| (4) - Oxygen (1 unit / Cost: 1)                               |")
		fmt.Println("| (5) - Repair Parts (1 unit / Cost: 3)                         |")
		fmt.Println("| (6) - Materials (1 unit / Cost: 4)                            |")
		fmt.Println("| (7) - Supplies (1 unit / Cost: 3)                             |")
		fmt.Println("| (8) - Passenger Compartment (25 passengers / Cost: 10)        |")

		switch user.Input() {
		case "1":
			gs.Fuel += 1
			gs.Budget -= 5
		case "2":
			gs.Food += 1
			gs.Budget -= 2
		case "3":
			gs.Water += 1
			gs.Budget -= 1
		case "4":
			gs.Oxygen += 1
			gs.Budget -= 1
		case "5":
			gs.RepairParts += 1
			gs.Budget -= 3
		case "6":
			gs.Materials += 1
			gs.Budget -= 4
		case "7":
			gs.Supplies += 1
			gs.Budget -= 3
		case "8":
			gs.Passengers += 25
			gs.Budget -= 10
		case "default":
			return
		}

		if gs.Budget <= 0 {
			return
		}
	}
}
