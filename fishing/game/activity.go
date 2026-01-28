package game

import (
	"home/gamecommon"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//Holds the logic for the fishing activity

type Activity struct{}

func init() {
	frame = 0
	goodframes = 0
	fishx, fishy = rand.Intn(ScreenWidth), rand.Intn(ScreenHeight)
	circle = gamecommon.NewCircle(ScreenWidth/2, ScreenHeight/2, 100)
}

var (
	fish         *Fish
	fishx, fishy int
	circle       gamecommon.Shape
	frame        int
	goodframes   int
)

func (activity *Activity) Draw(screen *ebiten.Image) {
	//Draw the circle
	geom := circle.Geometry.(gamecommon.Circle)
	vector.FillCircle(screen, float32(circle.X), float32(circle.Y),
		float32(geom.Radius), color.RGBA{0, 0, 255, 255}, true)

	//Draw the fish
	vector.FillRect(screen, float32(fishx), float32(fishy), float32(20), float32(10), color.Black, true)
}

func (activity *Activity) Update(g *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.mode = Title
		restart()
	}

	dir, ok := g.input.Dir()
	if ok {
		dx, dy := dir.Vector()
		fishx += dx * 50
		fishy += dy * 50
	}

	if frame%10 == 0 {
		dx, dy := gamecommon.RandomDir().Vector()
		fishx += 5 * dx
		fishy += 5 * dy
	}
	if circle.Contains(fishx, fishy) {
		goodframes++
	}
	if goodframes > 200 {
		g.catchFish()
		g.mode = Animation
		restart()
	}
	if frame > 1000 {
		restart()
	}
	return nil
}

func restart() {
	frame = 0
	goodframes = 0
	fishx, fishy = rand.Intn(ScreenWidth), rand.Intn(ScreenHeight)
}

func (g *Game) catchFish() {
	k := rand.Intn(len(fishes))
	for fish := range fishes {
		if k == 0 {
			held, found := g.save.Fish[fish]
			if !found {
				g.save.Fish[fish] = 1
			} else {
				g.save.Fish[fish] = held + 1
			}
			gamecommon.SaveGame(g.save)
			return
		}
		k--
	}
	panic("unreachable")
}
