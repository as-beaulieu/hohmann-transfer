package transferPhase

import "hohmannTransfer/src/game"

func Run(gs *game.State) {
	transferBurn(gs)
	inFlightEvent(gs)
	inFlightEvent(gs)
}
