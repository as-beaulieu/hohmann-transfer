package marsPhase

import (
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func Run(gs *game.State) {
	decelerationBurn(gs)
	enterOrbit(gs)
	landing(gs)
	graphics.ShowScore(gs.Score())
	user.Pause()
}
