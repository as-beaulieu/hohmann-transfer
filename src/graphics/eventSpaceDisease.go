package graphics

import "fmt"

func SpaceDiseaseEvent() {
	fmt.Println("|                                                                                |")
	fmt.Println("|  There is a space disease going around the ship                                |")
	fmt.Println("|                                                                                |")
}

func SpaceDiseaseCrewIll(member string) {
	fmt.Println("|                                                                                |")
	fmt.Printf(
		"  Your Mission %v Officer has fallen ill                                \n",
		member,
	)
	fmt.Println("  What do you want to do?")
	fmt.Println("  (1) - Open up the colony's medical supplies. (Requires 10 Supplies)  ")
	fmt.Println("  (2) - Trust your Medical Officer to create a cure.  ")
	fmt.Println("  (3) - Do nothing. It will be fine.")
	fmt.Println("|                                                                                |")
}
