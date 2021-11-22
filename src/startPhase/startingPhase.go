package startPhase

import (
	"hohmannTransfer/src/game"
)

func Run(gs *game.State) {
	missionSelect(gs)
	hireCrew(gs)
	procurement(gs)
}
