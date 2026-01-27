package game

import (
	"bytes"
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
	input = gamecommon.NewInput()
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
	input                  *gamecommon.Input
)

type Save struct {
	name  string
	fish  map[string]int
	money int
}

var save *Save
var shop *Shop
var titlePage *TitlePage

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
	shop = &Shop{
		selectedFish: "",
		selectedItem: "",
		quantity:     0,
	}

	titlePage = &TitlePage{}

	g := &Game{
		input: input,
		save:  save,
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (shop *Shop) increaseQuantity() {
	shop.quantity++
}

func (shop *Shop) dereaseQuantity() {
	if shop.quantity > 0 {
		shop.quantity--
	}
}

func (shop *Shop) resetSelection() {
	shop.selectedFish = ""
	shop.selectedItem = ""
	shop.quantity = 0
}

func (shop *Shop) makePurchase() {
	if shop.ShopMode == Buying {
		price := fishes[shop.selectedFish].Price
		totalCost := price * shop.quantity
		save.money -= totalCost
	}
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
		return nil
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

func (g *Game) drawFishing(screen *ebiten.Image) {

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
		shop.Draw(screen)
	case Fishing:
		g.drawFishing(screen)
	}
}
