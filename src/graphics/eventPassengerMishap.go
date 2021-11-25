package graphics

import "fmt"

func PassengerMishapEvent() {
	fmt.Println("|                                                                                |")
	fmt.Println("  A malfunction has occurred in one of the passenger pods.       ")
	fmt.Println("  Several passengers are trapped inside and the life support levels are dropping. ")
	fmt.Println("|                                                                                |")
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to restore life support and save the passengers         ")
	fmt.Println("  (2) - Use some of the colony provisions   (Requires 5 Repair Parts)                 ")
	fmt.Println("  (3) - Save what passengers you have outside the pod.                                ")
}
