package game

type State struct {
	Settings
	Mission
	Provisions
	Crew
}

type Settings struct {
	OperatingSystem string
}

func (gs *State) UseFuel(amount int) {
	if amount > gs.Fuel {
		gs.MissionFailure = true
	}

	gs.Fuel -= amount
}
