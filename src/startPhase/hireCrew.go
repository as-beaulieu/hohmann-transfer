package startPhase

import (
	"fmt"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func hireCrew(gs *game.State) {
	graphics.ClearScreen(gs.OperatingSystem)
	fmt.Println("| Purchase your crew:                 |")
	fmt.Println("| (1) - Novice (Cost: 5)              |")
	fmt.Println("| (2) - Experts (Cost: 15)            |")

	switch user.Input() {
	case "1":
		fmt.Println("Hiring the novice crew!")
		gs.Budget -= 5
		gs.Crew.Commander.Skill = 25
		gs.Crew.Pilot.Skill = 25
		gs.Crew.Medical.Skill = 25
		gs.Crew.Specialist.Skill = 25
		gs.Crew.Engineer.Skill = 25
		gs.Crew.Manifest = []string{"Commander", "Pilot", "Medical", "Specialist", "Engineer"}
	case "2":
		fmt.Println("Hiring the expert crew!")
		gs.Budget -= 15
		gs.Crew.Commander.Skill = 75
		gs.Crew.Pilot.Skill = 75
		gs.Crew.Medical.Skill = 75
		gs.Crew.Specialist.Skill = 75
		gs.Crew.Engineer.Skill = 75
		gs.Crew.Manifest = []string{"Commander", "Pilot", "Medical", "Specialist", "Engineer"}
	case "default":
		fmt.Println("exit the game")
		return
	}
	fmt.Println("Remaining budget: ", &gs.Budget)

}
