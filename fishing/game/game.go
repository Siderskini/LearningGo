package game

import (
	"bytes"
	"encoding/gob"
	"home/gamecommon"
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
	gob.Register(save)
	input = gamecommon.NewInput()
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFaceSource = s
	fishingAnimationFrames, err = gamecommon.ToEbitenFrames("./fishing/resources/fishing.gif", 120)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	arcadeFaceSource       *text.GoTextFaceSource
	fishingAnimationFrames []*ebiten.Image
	framecounter           int
	input                  *gamecommon.Input
)

type Save struct {
	Name      string
	Fish      map[string]int
	Inventory map[string]int
	Money     int
}

var save *Save
var shop *Shop
var titlePage *TitlePage
var activity *Activity

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	save, err := gamecommon.LoadGame(save)
	if err != nil {
		if os.IsNotExist(err) {
			save = &Save{
				Name:      "Sidd",
				Fish:      make(map[string]int),
				Inventory: make(map[string]int),
				Money:     0,
			}
		}
	}
	shop = &Shop{
		selectedFish:      "",
		selectedItem:      "",
		quantity:          0,
		displayedQuantity: "0",
	}

	titlePage = &TitlePage{}
	activity = &Activity{}

	g := &Game{
		input: input,
		save:  save.(*Save),
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
	switch g.mode {
	case Title:
		return titlePage.Update(g)
	case Animation:
		return nil
	case Shopping:
		return shop.Update(g)
	case Fishing:
		return activity.Update(g)
	}
	return nil
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

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	switch g.mode {
	case Title:
		titlePage.Draw(screen)
	case Animation:
		g.drawAnimation(screen)
	case Shopping:
		shop.Draw(g, screen)
	case Fishing:
		activity.Draw(screen)
	}
}
