package game

const (
	//Provisions scores

	FUEL_SCORE         = 1
	FOOD_SCORE         = 1
	WATER_SCORE        = 1
	OXYGEN_SCORE       = 1
	PASSENGER_SCORE    = 3
	REPAIR_PARTS_SCORE = 2
	MATERIALS_SCORE    = 2
	SUPPLIES_SCORE     = 2

	//Successful objectives

	SUCCESSFUL_SCORE_COLONY    = 100
	SUCCESSFUL_SCORE_SATELLITE = 25
	MISSION_CREW_ALIVE         = 50

	DIFFICULTY = 1
)

func (gs *State) Score() int {
	var score int

	score += gs.Budget

	score += gs.Fuel * FUEL_SCORE
	score += gs.Food * FOOD_SCORE
	score += gs.Water * WATER_SCORE
	score += gs.Oxygen * OXYGEN_SCORE
	score += gs.Passengers * PASSENGER_SCORE
	score += gs.RepairParts * REPAIR_PARTS_SCORE
	score += gs.Materials * MATERIALS_SCORE
	score += gs.Supplies * SUPPLIES_SCORE

	if gs.SuccessfulMission {
		switch gs.MissionType {
		case MISSION_TYPE_COLONY:
			score += SUCCESSFUL_SCORE_COLONY
		case MISSION_TYPE_SATELLITE:
			score += SUCCESSFUL_SCORE_SATELLITE
		default:
		}
	}

	if gs.MissionCrewAlive {
		score += MISSION_CREW_ALIVE
	}

	return score * DIFFICULTY
}
