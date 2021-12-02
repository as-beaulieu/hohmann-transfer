package transferPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func transferBurn(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.UseFuel(5)
	gs.StatusScreen()
	fmt.Println("|                                                                                |")
	fmt.Println("|  You successfully leave Earth orbit!                                           |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	fmt.Println("|                                                                                |")
	user.Pause()

	if gs.MissionFailure {
		fmt.Println("|                                                                                |")
		fmt.Println("|        You did not have enough fuel for the transfer!                          |")
		fmt.Println("|        Your rocket orbits high over the earth forever                          |")
		fmt.Println("|                                                                                |")
		user.Pause()
		graphics.GameOver()
		user.Pause()
	}
}
