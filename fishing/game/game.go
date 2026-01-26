package game

import (
	"bytes"
	"home/gamecommon"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth   = 420
	ScreenHeight  = 600
	boardSize     = 3
	titleFontSize = fontSize * 1.25
	fontSize      = 24
	smallFontSize = fontSize / 2
)

// Game represents a game state.
type Game struct {
	mode       Mode
	input      *gamecommon.Input
	boardImage *ebiten.Image
	save       *Save
}

type Mode int

const (
	Title Mode = iota
	Animation
	Shopping
	Fishing
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFaceSource = s

	fishingAnimationFrames, err = gamecommon.ToEbitenFrames("./fishing/resources/fishing.gif", 600)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	arcadeFaceSource       *text.GoTextFaceSource
	fishingAnimationFrames []*ebiten.Image
	framecounter           int
)

type Save struct {
	name  string
	fish  map[string]int
	money int
}

var save *Save

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	err := gamecommon.LoadGame(save)
	if err != nil {
		if os.IsNotExist(err) {
			save = &Save{
				name:  "Sidd",
				fish:  make(map[string]int),
				money: 0,
			}
		}
	}
	g := &Game{
		input: gamecommon.NewInput(),
		save:  save,
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()
	_, pressed := g.input.PressedTile()
	if pressed {
		g.mode = Animation
	}
	return nil
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	titleTexts := "FISHING GAME"
	texts := "\n\n\n\n\n\nPRESS SPACE KEY\n\nOR A/B BUTTON\n\nOR TOUCH SCREEN"
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, titleTexts, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   titleFontSize,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = fontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, texts, &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   fontSize,
	}, op)

	const msg = "Fishing Game by Sidd Viswanathan"

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

func (g *Game) drawAnimation(screen *ebiten.Image) {
	//Draw animation content here
	if framecounter >= len(fishingAnimationFrames) {
		framecounter = 0
		g.mode = Title
		return
	}
	currentImage := fishingAnimationFrames[framecounter]
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(currentImage, op)
	framecounter++
}

func (g *Game) drawShopping(screen *ebiten.Image) {

}

func (g *Game) drawFishing(screen *ebiten.Image) {

}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	switch g.mode {
	case Title:
		g.drawTitle(screen)
	case Animation:
		g.drawAnimation(screen)
	case Shopping:
		g.drawShopping(screen)
	case Fishing:
		g.drawFishing(screen)
	}
}
