package fishing

import (
	"home/fishing/game"
	"home/gamecommon"

	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	g := gamecommon.TryPanic(game.NewGame())
	mobile.SetGame(g)
}

// Dummy is a required exported function to force gomobile to compile the package.
func Dummy() {}
