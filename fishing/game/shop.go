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

var (
	shopButtonBackGroundColor = color.RGBA{0, 0, 255, 255}
	listenForQuantity         = false
)

// shopping buttons
var (
	buyButton      gamecommon.Button
	sellButton     gamecommon.Button
	fishButtons    []*gamecommon.Button
	storeButtons   []*gamecommon.Button
	leaveButton    gamecommon.Button
	plusButton     gamecommon.Button
	minusButton    gamecommon.Button
	confirmButton  gamecommon.Button
	quantityButton gamecommon.Button
	buyMenu        *gamecommon.ScrollMenu
	sellMenu       *gamecommon.ScrollMenu
)

var storeItems = map[string]Item{
	"Bait":     Item{"Bait", 15},
	"Good Rod": Item{"Good Rod", 100},
}

func init() {
	for item := range storeItems {
		storeButtons = append(storeButtons, newStoreButton(item))
	}
	sellMenu = gamecommon.NewScrollMenu(fishButtons, "Sell Fish", ScreenWidth/2-200, 3*titleFontSize, 400, 300, 40, input)
	buyMenu = gamecommon.NewScrollMenu(storeButtons, "Buy Fish", ScreenWidth/2-200, 3*titleFontSize, 400, 300, 40, input)
	buyButton = gamecommon.NewRectangleButton("buy", ScreenWidth/2-200, 1*titleFontSize, 200-2, 50, "Buy", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	sellButton = gamecommon.NewRectangleButton("sell", ScreenWidth/2, 1*titleFontSize, 200, 50, "Sell", arcadeFaceSource, color.RGBA{255, 255, 0, 255}, color.RGBA{0, 0, 255, 255}, input)
	leaveButton = gamecommon.NewRectangleButton("leave", ScreenWidth/2-200, 16*titleFontSize, 100, 50, "Leave", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	plusButton = gamecommon.NewRectangleButton("plus", ScreenWidth/2-80, 14*titleFontSize, 60, 50, "+", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	minusButton = gamecommon.NewRectangleButton("minus", ScreenWidth/2-200, 14*titleFontSize, 60, 50, "-", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	confirmButton = gamecommon.NewRectangleButton("confirm", ScreenWidth/2, 14*titleFontSize, 200, 50, "Confirm", arcadeFaceSource, color.White, color.RGBA{0, 0, 255, 255}, input)
	quantityButton = gamecommon.NewRectangleButton("quantity", ScreenWidth/2-140, 14*titleFontSize, 60, 50, "0", arcadeFaceSource, color.White, backgroundColor, input)
}

func newStoreButton(itemName string) *gamecommon.Button {
	item := storeItems[itemName]
	buttonWidth := 30
	priceString := strconv.Itoa(item.Price)
	spaces := buttonWidth - len(priceString) - len(item.Name)
	text := item.Name + string(bytes.Repeat([]byte(" "), spaces)) + "$" + priceString
	btn := gamecommon.NewRectangleButton(item.Name, 0, 0, 0, 0, text, arcadeFaceSource, color.White, shopButtonBackGroundColor, input)
	return &btn
}

func newFishButton(fish string, value int) *gamecommon.Button {
	buttonWidth := 30
	priceString := strconv.Itoa(fishes[fish].Price)
	spaces := buttonWidth - len(priceString) - len(fish) - 3
	text := fish + string(bytes.Repeat([]byte(" "), spaces)) + "x" + strconv.Itoa(value) + " $" + priceString
	btn := gamecommon.NewRectangleButton(fish, 0, 0, 0, 0, text, arcadeFaceSource, color.White, shopButtonBackGroundColor, input)
	return &btn
}

func (shop *Shop) Draw(screen *ebiten.Image) {
	// Draw buttons and menus
	buyButton.Draw(screen)
	sellButton.Draw(screen)
	leaveButton.Draw(screen)
	plusButton.Draw(screen)
	minusButton.Draw(screen)
	confirmButton.Draw(screen)
	quantityButton.Text = strconv.Itoa(shop.quantity)
	quantityButton.Draw(screen)
	if shop.ShopMode == Buying {
		buyMenu.Draw(screen)
	} else {
		sellMenu.Draw(screen)
	}

	// Draw the player's current money
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2+20, 16*titleFontSize+15)
	op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 0, 255})
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, "Money: "+strconv.Itoa(save.money), &text.GoTextFace{
		Source: arcadeFaceSource,
		Size:   24,
	}, op)

	// Draw the results of the transaction
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2+150, 16*titleFontSize+15)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	if shop.ShopMode == Buying {
		op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
		text.Draw(screen, "-"+strconv.Itoa(storeItems[shop.selectedItem].Price*shop.quantity), &text.GoTextFace{
			Source: arcadeFaceSource,
			Size:   24,
		}, op)
	} else {
		op.ColorScale.ScaleWithColor(color.RGBA{0, 255, 0, 255})
		text.Draw(screen, "+"+strconv.Itoa(fishes[shop.selectedFish].Price*shop.quantity), &text.GoTextFace{
			Source: arcadeFaceSource,
			Size:   24,
		}, op)
	}
}

func (shop *Shop) Update(g *Game) error {
	sellMenu = gamecommon.NewScrollMenu(fishButtons, "Sell Fish", ScreenWidth/2-200, 3*titleFontSize, 400, 300, 40, input)
	if shop.ShopMode == Selling {
		newFish := sellMenu.HandleInput()
		if newFish != "" && newFish != shop.selectedFish {
			shop.resetSelection()
			shop.selectedFish = newFish
			return nil
		}
	} else {
		newItem := buyMenu.HandleInput()
		if newItem != shop.selectedItem {
			shop.resetSelection()
			shop.selectedItem = newItem
			return nil
		}
	}

	pressed := buyButton.IsPressed()
	if pressed && shop.ShopMode == Selling {
		toggleShopMode(shop)
		return nil
	}
	pressed = sellButton.IsPressed()
	if pressed && shop.ShopMode == Buying {
		toggleShopMode(shop)
		return nil
	}
	pressed = plusButton.IsPressed()
	if pressed {
		shop.increaseQuantity()
		return nil
	}
	pressed = minusButton.IsPressed()
	if pressed {
		shop.dereaseQuantity()
		return nil
	}
	pressed = leaveButton.IsPressed()
	if pressed {
		g.mode = Title
		shop.resetSelection()
		return nil
	}
	pressed = confirmButton.IsPressed()
	if pressed {
		shop.makePurchase()
		return nil
	}
	pressed = quantityButton.IsPressed()
	if pressed {
		listenForQuantity = true
	}
	if listenForQuantity {
		s := input.TextInput()
		if s != "" {
			quantity, err := strconv.Atoi(s)
			if err == nil {
				shop.quantity = quantity
			}
			listenForQuantity = false
			return err
		}
	}
	fishButtons = []*gamecommon.Button{}
	for fish, value := range g.save.fish {
		fishButtons = append(fishButtons, newFishButton(fish, value))
	}
	return nil
}

func toggleShopMode(shop *Shop) {
	if shop.ShopMode == Buying {
		shop.ShopMode = Selling
		sellButton.Color = color.RGBA{255, 255, 0, 255}
		buyButton.Color = color.White
	} else {
		shop.ShopMode = Buying
		sellButton.Color = color.White
		buyButton.Color = color.RGBA{255, 255, 0, 255}
	}
	shop.resetSelection()
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
		price := storeItems[shop.selectedItem].Price
		totalCost := price * shop.quantity
		save.money -= totalCost
	} else {
		price := fishes[shop.selectedFish].Price
		totalCost := price * shop.quantity
		save.money += totalCost
	}
}
