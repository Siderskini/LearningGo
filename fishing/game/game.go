package game

import (
	"bytes"
	"home/gamecommon"
	"image/color"
	"log"
	"os"
	"strconv"

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

type Shop struct {
	ShopMode     ShopMode
	selectedFish string
	selectedItem string
	quantity     int
}

type Mode int

const (
	Title Mode = iota
	Animation
	Shopping
	Fishing
)

type ShopMode bool

const (
	Buying  ShopMode = true
	Selling ShopMode = false
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

	fishingButton = gamecommon.NewRectangleButton("fish", ScreenWidth/2-100, 5*titleFontSize, 200, 50, "Click to Start", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	shoppingButton = gamecommon.NewRectangleButton("shop", ScreenWidth/2-100, 7*titleFontSize, 200, 50, "Click to Shop", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	quitButton = gamecommon.NewRectangleButton("quit", ScreenWidth/2-100, 9*titleFontSize, 200, 50, "Click to Quit", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	for fish := range fishes {
		buttonText := fishButtonText(fish)
		fishButton := gamecommon.NewRectangleButton(fish, 0, 0, 0, 0, buttonText, arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
		fishButtons = append(fishButtons, fishButton)
	}
	sellMenu = gamecommon.NewScrollMenu(fishButtons, "Sell Fish", ScreenWidth/2-200, 3*titleFontSize, 400, 300, 40, input)
	buyButton = gamecommon.NewRectangleButton("buy", ScreenWidth/2-200, 1*titleFontSize, 200, 50, "Buy", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	sellButton = gamecommon.NewRectangleButton("sell", ScreenWidth/2, 1*titleFontSize, 200, 50, "Sell", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	leaveButton = gamecommon.NewRectangleButton("leave", ScreenWidth/2-200, 16*titleFontSize, 100, 50, "Leave", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	plusButton = gamecommon.NewRectangleButton("plus", ScreenWidth/2-80, 14*titleFontSize, 60, 50, "+", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	minusButton = gamecommon.NewRectangleButton("minus", ScreenWidth/2-200, 14*titleFontSize, 60, 50, "-", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	confirmButton = gamecommon.NewRectangleButton("confirm", ScreenWidth/2, 14*titleFontSize, 200, 50, "Confirm", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
}

func fishButtonText(fish string) string {
	buttonWidth := 30
	priceString := strconv.Itoa(fishes[fish].Price)
	spaces := buttonWidth - len(priceString) - len(fish)
	return fish + string(bytes.Repeat([]byte(" "), spaces)) + "$" + priceString
}

var (
	arcadeFaceSource       *text.GoTextFaceSource
	fishingAnimationFrames []*ebiten.Image
	framecounter           int
	input                  *gamecommon.Input

	//title buttons
	fishingButton  gamecommon.Button
	shoppingButton gamecommon.Button
	quitButton     gamecommon.Button

	//shopping buttons
	buyButton     gamecommon.Button
	sellButton    gamecommon.Button
	fishButtons   []gamecommon.Button
	storeButtons  []gamecommon.Button
	leaveButton   gamecommon.Button
	plusButton    gamecommon.Button
	minusButton   gamecommon.Button
	confirmButton gamecommon.Button
	buyMenu       *gamecommon.ScrollMenu
	sellMenu      *gamecommon.ScrollMenu

	//fishing buttons
)

type Save struct {
	name  string
	fish  map[string]int
	money int
}

var save *Save
var shop *Shop

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
		pressed := fishingButton.IsPressed()
		if pressed {
			g.mode = Animation
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
	case Animation:
		return nil
	case Shopping:
		if shop.ShopMode == Selling {
			sellMenu.HandleInput()
		}

		pressed := buyButton.IsPressed()
		if pressed {
			shop.ShopMode = Buying
			shop.resetSelection()
		}
		pressed = sellButton.IsPressed()
		if pressed {
			shop.ShopMode = Selling
			shop.resetSelection()
		}
		pressed = plusButton.IsPressed()
		if pressed {
			g.mode = Shopping
			shop.increaseQuantity()
		}
		pressed = minusButton.IsPressed()
		if pressed {
			g.mode = Shopping
			shop.dereaseQuantity()
		}
		pressed = leaveButton.IsPressed()
		if pressed {
			g.mode = Title
			shop.resetSelection()
		}
		pressed = confirmButton.IsPressed()
		if pressed {
			g.mode = Shopping
		}
		return nil
	case Fishing:
		return nil
	}
	return nil
}

func (g *Game) drawTitle(screen *ebiten.Image) {
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
	buyButton.Draw(screen)
	sellButton.Draw(screen)
	leaveButton.Draw(screen)
	plusButton.Draw(screen)
	minusButton.Draw(screen)
	confirmButton.Draw(screen)
	sellMenu.Draw(screen)

	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2-110, 14*titleFontSize+10)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, strconv.Itoa(shop.quantity), &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   24,
	}, op)
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
