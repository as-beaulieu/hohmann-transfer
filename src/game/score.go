package game

var (
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

	COLONY_SUCCESSFUL_SCORE    = 100
	SATELLITE_SUCCESSFUL_SCORE = 25
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

	return score
}
