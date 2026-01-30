package fishing

import (
	"home/fishing/game"
	"home/gamecommon"

	"github.com/hajimehoshi/ebiten/v2"
)

func StartGame() {
	g := gamecommon.TryPanic(game.NewGame())
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Fishing Game")
	gamecommon.TryPanic("", ebiten.RunGame(g))
}
