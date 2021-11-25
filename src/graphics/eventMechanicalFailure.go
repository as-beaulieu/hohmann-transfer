package graphics

import "fmt"

func MechanicalFailureEvent() {
	fmt.Println("|                                                                                |")
	fmt.Println("  A malfunction has occurred.       ")
	fmt.Println("  If not repaired it could jeopardize the entire mission. ")
	fmt.Println("|                                                                                |")
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to fix the issue         ")
	fmt.Println("  (2) - Use some of the colony provisions   (Requires 5 Repair Parts)                 ")
	fmt.Println("  (3) - Continue the mission with the broken parts.                                ")
}
