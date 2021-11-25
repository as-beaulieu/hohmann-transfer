package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) spaceDiseaseEvent() {
	graphics.SpaceDiseaseEvent()
	user.Pause()

	member, err := gs.Crew.WhoToKill()
	if err != nil {
		gs.MissionFailure = true
		return
	}
	graphics.SpaceDiseaseCrewIll(member)
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
		if gs.Medical.TestSkill(50) {
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
		member,
		disease,
	)
	fmt.Println("|                                                                                |")

	gs.Crew.KillCrewMember(member)
}

func spaceDisease() string {
	var diseases = []string{"Space Dysentery", "Space Snakebite", "Space Cholera", "Space Typhoid"}
	i := Random(4)
	return diseases[i]
}
