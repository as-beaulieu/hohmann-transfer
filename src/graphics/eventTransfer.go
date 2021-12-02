package graphics

import "fmt"

func FoodSpoilageEvent() {
	borderedFiller()
	fmt.Println("  Your mission crew reports a malfunction in the ship's refrigerators       ")
	fmt.Println("  All of your mission's food stores are at risk of spoiling       ")
	borderedFiller()
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to try to patch the refrigerators                ")
	fmt.Println("  (2) - Use some of the colony supplies   (Requires 5 Repair Parts)                   ")
	fmt.Println("  (3) - Move whatever food that will fit in the working refrigerators                 ")
}

func MechanicalFailureEvent() {
	borderedFiller()
	fmt.Println("  A malfunction has occurred.       ")
	fmt.Println("  If not repaired it could jeopardize the entire mission. ")
	borderedFiller()
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to fix the issue         ")
	fmt.Println("  (2) - Use some of the colony provisions   (Requires 5 Repair Parts)                 ")
	fmt.Println("  (3) - Continue the mission with the broken parts.                                ")
}

func PassengerMishapEvent() {
	borderedFiller()
	fmt.Println("  A malfunction has occurred in one of the passenger pods.       ")
	fmt.Println("  Several passengers are trapped inside and the life support levels are dropping. ")
	borderedFiller()
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to restore life support and save the passengers         ")
	fmt.Println("  (2) - Use some of the colony provisions   (Requires 5 Repair Parts)                 ")
	fmt.Println("  (3) - Save what passengers you have outside the pod.                                ")
}

func SpaceDiseaseEvent() {
	borderedFiller()
	fmt.Println("|  There is a space disease going around the ship                                |")
	borderedFiller()
}

func SpaceDiseaseCrewIll(member string) {
	borderedFiller()
	fmt.Printf(
		"  Your Mission %v Officer has fallen ill                                \n",
		member,
	)
	fmt.Println("  What do you want to do?")
	fmt.Println("  (1) - Open up the colony's medical supplies. (Requires 10 Supplies)  ")
	fmt.Println("  (2) - Trust your Medical Officer to create a cure.  ")
	fmt.Println("  (3) - Do nothing. It will be fine.")
	borderedFiller()
}

func SpaceSnakeEvent() {
	borderedFiller()
	fmt.Println("  One of the passengers had snuck their snake in their personal belongings.       ")
	fmt.Println("  The snake has escaped and roaming loose around the ship.       ")
	borderedFiller()
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your mission specialist to capture the snake                ")
	fmt.Println("  (2) - Tell the passengers to clean up their own mess                      ")
	fmt.Println("  (3) - Do nothing. It will be the ship mascot.                             ")
}
