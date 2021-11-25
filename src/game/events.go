package game

import (
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
	"math/rand"
	"time"
)

func Random(ceiling int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)
	return r1.Intn(ceiling-1) + 1
}

func (gs *State) DetermineTransferEvent() {
	result := Random(100)
	if result <= 10 {
		graphics.UnremarkableEvent()
	}
	if result >= 11 && result <= 20 {
		gs.foodSpoilageEvent()
	}
	if result >= 21 && result <= 30 {
		gs.passengerMishapEvent()
	}
	if result >= 31 && result <= 50 {
		gs.mechanicalFailureEvent()
	}
	if result >= 51 && result <= 60 {
		gs.spaceDiseaseEvent()
	}
	if result >= 61 && result <= 80 {
		graphics.UnremarkableEvent()
	}
	if result >= 81 && result <= 100 {
		gs.spaceSnakeEvent()
	}
	user.Pause()
}
