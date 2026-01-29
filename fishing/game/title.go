package game

import (
	"home/gamecommon"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitlePage struct{}

const titleFontSize = fontSize * 1.25

// title buttons
var (
	fishingButton  gamecommon.Button
	shoppingButton gamecommon.Button
	quitButton     gamecommon.Button
)

func init() {
	fishingButton = gamecommon.NewRectangleButton("fish", ScreenWidth/2-100, 5*titleFontSize, 200, 50, "Click to Start", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	shoppingButton = gamecommon.NewRectangleButton("shop", ScreenWidth/2-100, 7*titleFontSize, 200, 50, "Click to Shop", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	quitButton = gamecommon.NewRectangleButton("quit", ScreenWidth/2-100, 9*titleFontSize, 200, 50, "Click to Quit", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
}

func (title *TitlePage) Draw(screen *ebiten.Image) {
	fishingButton.Draw(screen)
	shoppingButton.Draw(screen)
	quitButton.Draw(screen)

	titleTexts := "FISHING GAME"
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, titleTexts, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   titleFontSize,
	}, op)

	const msg = "Fishing Game by Sidd Viswanathan and Olga Nam"

	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight-smallFontSize/2)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = smallFontSize
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignEnd
	text.Draw(screen, msg, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   smallFontSize,
	}, op)
}

func (title *TitlePage) Update(g *Game) error {
	pressed := fishingButton.IsPressed()
	if pressed {
		g.catchFish()
		g.mode = Fishing
	}
	pressed = shoppingButton.IsPressed()
	if pressed {
		shop.ShopMode = Selling
		g.mode = Shopping
	}
	pressed = quitButton.IsPressed()
	if pressed {
		return ebiten.Termination
	}
	return nil
}
