package earthPhase

import (
	"hohmannTransfer/src/game"
)

func Run(gs *game.State) {
	earthLaunch(gs)
	earthOrbit(gs)
}
