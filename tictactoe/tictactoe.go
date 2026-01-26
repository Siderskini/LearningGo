package tictactoe

import (
	"home/tictactoe/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func StartGame() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Tic Tac Toe")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
