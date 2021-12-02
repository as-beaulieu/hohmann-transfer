package game

import "runtime"

type State struct {
	Settings
	Mission
	Provisions
	Crew
	ExitGame bool
}

type Settings struct {
	OperatingSystem string
}

func NewGameState() *State {
	return &State{
		Mission:    Mission{},
		Settings:   Settings{OperatingSystem: runtime.GOOS},
		Provisions: Provisions{},
		Crew:       Crew{},
	}
}

func (gs *State) UseFuel(amount int) {
	if amount > gs.Fuel {
		gs.MissionFailure = true
	}

	gs.Fuel -= amount
}
