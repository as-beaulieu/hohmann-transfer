package startPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/user"
)

func missionSelect(gs *game.State) {
	fmt.Println("|                                                                                |")
	fmt.Println("|  Choose your mission to send to Mars!                                          |")
	fmt.Println("|  (1) - Send a colony                                                           |")
	fmt.Println("|  (2) - Send a satellite (Not Implemented Yet)                                  |")
	fmt.Println("|                                                                                |")

	switch user.Input() {
	case "1":
		fmt.Println("Sending a Colony!")
		gs.MissionType = game.MISSION_TYPE_COLONY
		gs.Passengers += 50
		gs.Budget = 200
	case "2":
		fmt.Println("Sending a Satellite!")
		gs.MissionType = game.MISSION_TYPE_SATELLITE
		gs.Budget = 75
	case "default":
		fmt.Println("exit the game")
		return
	}

}
