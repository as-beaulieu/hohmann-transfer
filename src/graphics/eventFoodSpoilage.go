package graphics

import "fmt"

func FoodSpoilageEvent() {
	fmt.Println("|                                                                                |")
	fmt.Println("  Your mission crew reports a malfunction in the ship's refrigerators       ")
	fmt.Println("  All of your mission's food stores are at risk of spoiling       ")
	fmt.Println("|                                                                                |")
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Send in your engineer to try to patch the refrigerators                ")
	fmt.Println("  (2) - Use some of the colony supplies   (Requires 5 Repair Parts)                   ")
	fmt.Println("  (3) - Move whatever food that will fit in the working refrigerators                 ")
}
