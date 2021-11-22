package graphics

import "runtime"

type GraphicSettings struct {
	OperatingSystem string
}

func Setup() *GraphicSettings{
	return &GraphicSettings{
		OperatingSystem: runtime.GOOS,
	}
}