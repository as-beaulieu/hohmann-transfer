package marsPhase

import (
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
)

func Run(gs *game.State) {
	decelerationBurn(gs)
	enterOrbit(gs)
	landing(gs)
	graphics.ShowScore(gs.Score())
}
