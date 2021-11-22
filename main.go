package main

import (
	"fmt"
	"hohmannTransfer/src/earthPhase"
	"hohmannTransfer/src/game"
	"hohmannTransfer/src/graphics"
	"hohmannTransfer/src/marsPhase"
	"hohmannTransfer/src/startPhase"
	"hohmannTransfer/src/transferPhase"
	"hohmannTransfer/src/user"
	"runtime"
)

func main() {
	gameState := game.State{
		Settings: game.Settings{
			OperatingSystem: runtime.GOOS,
		},
	}

	startScreen(&gameState)
}

func startScreen(gameState *game.State) {
	for {
		graphics.ClearScreen(gameState.OperatingSystem)
		graphics.StartScreen()

		switch user.Input() {
		case "s":
			fmt.Println("Start the Game")
			startGame(gameState)
		case "o":
			fmt.Println("options")
			options(gameState)
		case "e":
			fmt.Println("exit the game")
			return
		}
	}
}

func startGame(gameState *game.State) {
	graphics.ClearScreen(gameState.OperatingSystem)
	startPhase.Run(gameState)
	earthPhase.Run(gameState)
	transferPhase.Run(gameState)
	marsPhase.Run(gameState)
}

func options(gameState *game.State) {
	graphics.ClearScreen(gameState.OperatingSystem)
	graphics.OptionsScreen()
	user.Input()
}
