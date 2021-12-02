package marsPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func decelerationBurn(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.UseFuel(5)
	gs.StatusScreen()
	fmt.Println("|                                                                                |")
	fmt.Println("|  You begin your deceleration burn for Mars orbit!                              |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	fmt.Println("|                                                                                |")
	user.Pause()

	if gs.MissionFailure {
		fmt.Println("|                                                                                |")
		fmt.Println("|        You did not have enough fuel for deceleration!                          |")
		fmt.Println("|        Your rocket slingshots past Mars and into deep space!                   |")
		fmt.Println("|                                                                                |")
		user.Pause()
		graphics.GameOver()
		user.Pause()
	}
}
