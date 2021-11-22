package game

import (
	"fmt"
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
		fmt.Println("Food Spoilage")
	}
	if result >= 21 && result <= 30 {
		fmt.Println("Passenger Mishap")
	}
	if result >= 31 && result <= 40 {
		fmt.Println("Mechanical Failure")
	}
	if result >= 41 && result <= 50 {
		gs.spaceDiseaseEvent()
	}
	if result >= 51 && result <= 80 {
		graphics.UnremarkableEvent()
	}
	if result >= 81 && result <= 100 {
		gs.spaceSnakeEvent()
	}
	user.Pause()
}
