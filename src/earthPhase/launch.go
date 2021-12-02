package earthPhase

import (
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func earthLaunch(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	gs.UseFuel(5)
	gs.StatusScreen()
	graphics.EarthLaunch()
	user.Pause()

	gs.DetermineEarthManeuverEvent()

	if gs.MissionFailure {
		graphics.EarthLaunchFailure()
		user.Pause()
		graphics.GameOver()
		user.Pause()
	}
}
