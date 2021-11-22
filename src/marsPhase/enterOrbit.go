package marsPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func enterOrbit(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.UseFuel(5)
	gs.StatusScreen()
	fmt.Println("|                                                                                |")
	fmt.Println("|  You begin your orbital burn for Mars!                                         |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	fmt.Println("|                                                                                |")
	user.Pause()

	if gs.MissionFailure {
		fmt.Println("|                                                                                |")
		fmt.Println("|        You did not have enough fuel for orbit!                                 |")
		fmt.Println("|        Your rocket floats high over Mars forever!                              |")
		fmt.Println("|                                                                                |")
		user.Pause()
		graphics.GameOver()
	}
}
