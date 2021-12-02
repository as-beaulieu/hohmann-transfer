package earthPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func earthOrbit(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.UseFuel(5)
	gs.StatusScreen()
	fmt.Println("|                                                                                |")
	fmt.Println("|  You successfully enter orbit                                                  |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	fmt.Println("|                                                                                |")
	user.Pause()

	if gs.MissionFailure {
		graphics.EarthOrbitFailure()
		user.Pause()
		graphics.GameOver()
		user.Pause()
	}
}
