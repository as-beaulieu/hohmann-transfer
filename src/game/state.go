package game

type State struct {
	Settings
	Mission        string
	MissionFailure bool
	Budget         int
	Provisions
	Crew
}

type Settings struct {
	OperatingSystem string
}

type Provisions struct {
	Fuel        int
	Passengers  int
	Food        int
	Water       int
	Oxygen      int
	RepairParts int
	Materials   int
	Supplies    int
}

func (gs *State) UseFuel(amount int) {
	if amount > gs.Fuel {
		gs.MissionFailure = true
	}

	gs.Fuel -= amount
}
