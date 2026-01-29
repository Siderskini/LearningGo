package game

import (
	"home/gamecommon"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Sets up a new save
type Initial struct{}

func init() {
	listeningForName = false
	nameButton = gamecommon.NewRectangleButton("Name", ScreenWidth/2-150, ScreenHeight/2+150, 300, 50, "Enter name here", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
}

var nameButton gamecommon.Button
var listeningForName bool

func (initial *Initial) Draw(screen *ebiten.Image) {
	titleTexts := "WELCOME!\n\nWHAT IS\nYOUR NAME?"
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, titleTexts, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   titleFontSize,
	}, op)

	nameButton.Draw(screen)
}

func (initial *Initial) Update(g *Game) error {
	if nameButton.IsPressed() {
		listeningForName = true
		input.TextInputStart()
	}
	if listeningForName {
		temp, done := input.TextInput()
		if done {
			g.save = &Save{
				Name:      temp,
				Fish:      make(map[string]int),
				Inventory: make(map[string]int),
				Money:     0,
			}
			g.mode = Title
			gamecommon.SaveGame(g.save)
		}
		nameButton.Text = temp
	}
	return nil
}
