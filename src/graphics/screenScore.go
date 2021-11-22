package graphics

import (
	"fmt"
	"os"
)

func ShowScore(score int) {
	fmt.Println("|                                                                                |")
	fmt.Println("|                                                                                |")
	fmt.Printf(
		"        Your final score: %v                                                  \n",
		score,
	)
	fmt.Println("|                                                                                |")
	fmt.Println("|                                                                                |")
	os.Exit(0)
}
