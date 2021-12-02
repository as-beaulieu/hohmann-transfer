package graphics

import (
	"fmt"
)

func GameOver() {
	borderedFiller()
	borderedFiller()
	fmt.Println("|        Game Over!                                                              |")
	fmt.Println("|        You did not survive the Hohmann Transfer                                |")
	borderedFiller()
	borderedFiller()
}
