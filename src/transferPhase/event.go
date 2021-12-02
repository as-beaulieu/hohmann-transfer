package transferPhase

import (
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func inFlightEvent(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.StatusScreen()

	gs.DetermineTransferEvent()

	if gs.MissionFailure {
		user.Pause()
		graphics.GameOver()
		user.Pause()
	}
}
