package fishing

import (
	"home/fishing/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func StartGame() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Fishing Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
