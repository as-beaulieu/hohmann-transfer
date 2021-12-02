package graphics

import (
	"fmt"
)

func ShowScore(score int) {
	borderedFiller()
	borderedFiller()
	fmt.Printf(
		"        Your final score: %v                                                  \n",
		score,
	)
	borderedFiller()
	borderedFiller()
}
