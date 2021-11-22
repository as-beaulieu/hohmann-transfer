package game

import "fmt"

func (gs *State) StatusScreen() {
	fmt.Println("|Mission Status:                                                                 |")
	fmt.Printf(
		"   fuel: %v    passengers: %v     food: %v     water: %v     oxygen: %v        \n",
		gs.Fuel,
		gs.Passengers,
		gs.Food,
		gs.Water,
		gs.Oxygen,
	)
	fmt.Printf(
		"   repair parts: %v    materials: %v     supplies: %v     budget: %v           \n",
		gs.RepairParts,
		gs.Materials,
		gs.Supplies,
		gs.Budget,
	)
	fmt.Printf(
		"  Commander: %v   Pilot: %v   Medical: %v   Specialist: %v   Engineer: %v        \n",
		gs.Commander.Status(),
		gs.Pilot.Status(),
		gs.Medical.Status(),
		gs.Specialist.Status(),
		gs.Engineer.Status(),
	)
	fmt.Println("|                                                                                |")
}
