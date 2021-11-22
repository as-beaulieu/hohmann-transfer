package graphics

import (
	"fmt"
	"os"
)

func GameOver() {
	fmt.Println("|                                                                                |")
	fmt.Println("|                                                                                |")
	fmt.Println("|        Game Over!                                                              |")
	fmt.Println("|        You did not survive the Hohmann Transfer                                |")
	fmt.Println("|                                                                                |")
	fmt.Println("|                                                                                |")
	os.Exit(0)
}
