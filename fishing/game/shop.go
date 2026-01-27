package game

import (
	"bytes"
	"home/gamecommon"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Shop struct {
	ShopMode     ShopMode
	selectedFish string
	selectedItem string
	quantity     int
}

type ShopMode bool

const (
	Buying  ShopMode = true
	Selling ShopMode = false
)

// shopping buttons
var (
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
)

func init() {
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

func (*Shop) Draw(screen *ebiten.Image) {
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

func (*Shop) Update(g *Game) error {
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
