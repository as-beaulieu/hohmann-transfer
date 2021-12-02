package graphics

import (
	"fmt"
	"runtime"
)

type GraphicSettings struct {
	OperatingSystem string
}

func Setup() *GraphicSettings {
	return &GraphicSettings{
		OperatingSystem: runtime.GOOS,
	}
}

func borderedFiller() {
	fmt.Println("|                                                                                |")
}
