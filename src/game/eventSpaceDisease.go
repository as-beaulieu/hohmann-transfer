package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) spaceDiseaseEvent() {
	graphics.SpaceDiseaseEvent()
	user.Pause()

	i := Random(5)
	graphics.SpaceDiseaseCrewIll(gs.Manifest[i])
	switch user.Input() {
	case "1":
		if gs.Supplies >= 10 {
			fmt.Println("You use your mission's medical supplies")
			gs.Supplies -= 10
			return
		} else {
			fmt.Println("You do not have enough supplies!")
		}
	case "2":
		if !gs.Medical.Dead && gs.Medical.TestSkill(50) {
			fmt.Println("Your medical officer cured them")
			return
		}
	default:
		fmt.Println("You do nothing. What's the worse that could happen?")
	}

	disease := spaceDisease()
	fmt.Println("|                                                                                |")
	fmt.Printf(
		"  Your Mission %v Officer has died of %v                           \n",
		gs.Manifest[i],
		disease,
	)
	fmt.Println("|                                                                                |")

	switch gs.Manifest[i] {
	case "Commander":
		gs.Commander.Dead = true
	case "Pilot":
		gs.Pilot.Dead = true
	case "Medical":
		gs.Medical.Dead = true
	case "Specialist":
		gs.Specialist.Dead = true
	case "Engineer":
		gs.Engineer.Dead = true
	}

}

func spaceDisease() string {
	var diseases = []string{"Space Dysentery", "Space Snakebite", "Space Cholera", "Space Typhoid"}
	i := Random(4)
	return diseases[i]
}
