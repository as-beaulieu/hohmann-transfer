package graphics

import (
	"fmt"
	"os"
	"os/exec"
)

// ClearScreen executes the clear command per operating system
func ClearScreen(operatingSystem string) {
	fmt.Println(operatingSystem)
	fmt.Println("Clearing Screen!")
	fmt.Println("----------------")
	switch operatingSystem {
	case "linux":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("operating system not supported")
	}
}
