package game

import (
	"fmt"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/user"
)

func (gs *State) offCourseManeuver() {
	graphics.ManeuverOffCourse()

	switch user.Input() {
	case "1":
		if gs.Pilot.TestSkill(50) {
			fmt.Println("Your pilot was able to make a correction mid-maneuver.")
			gs.UseFuel(1)
			return
		}
		fmt.Println("The correction was not successful. Extra fuel was necessary to make a full correction.")
		gs.UseFuel(Random(5))
	case "2":
		if gs.Commander.TestSkill(75) {
			fmt.Println("Your commander made a new course to correct over the transfer to minimize fuel.")
			gs.UseFuel(Random(5))
			return
		}
		fmt.Println("The correction was not successful. Extra fuel was necessary to make a full correction.")
		gs.UseFuel(Random(10))
	default:
		fmt.Println("you do nothing")
	}

}
