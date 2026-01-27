package gamecommon

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScrollMenu struct {
	x, y, width, height, buttonheight int
	buttons                           []Button
	input                             *Input
	scrollMenuOffset                  int
}

/*
Create a new ScrollMenu with the given buttons and dimensions.
buttons: slice of Button objects to display in the menu
*/
func NewScrollMenu(buttons []Button, title string, x, y, width, height, buttonheight int, input *Input) *ScrollMenu {
	for i, button := range buttons {
		switch g := button.shape.geometry.(type) {
		case Rectangle:
			g.Width = width - 20
			g.Height = buttonheight - 2
			button.shape.geometry = g
		default:
			panic("ScrollMenu only supports Rectangle buttons")
		}
		button.shape.X = x
		buttons[i] = button
	}
	return &ScrollMenu{
		buttons:          buttons,
		x:                x,
		y:                y,
		width:            width,
		height:           height,
		buttonheight:     buttonheight,
		scrollMenuOffset: 0,
		input:            input,
	}
}

func (sm *ScrollMenu) Draw(screen *ebiten.Image) {
	menuImage := ebiten.NewImage(sm.width, sm.height)
	menuImage.Fill(color.Black)
	for i, button := range sm.buttons {
		button.shape.Y = sm.y + (i-sm.scrollMenuOffset-2)*sm.buttonheight
		button.Draw(menuImage)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sm.x), float64(sm.y))
	screen.DrawImage(menuImage, op)
}

func (sm *ScrollMenu) HandleInput() {
	for _, button := range sm.buttons {
		pressed := button.IsPressed()
		if pressed {
			// Handle button press logic here
		}
	}

	//Handle the scrolling
	dir, pressed := sm.input.Dir()
	if !pressed {
		return
	}
	switch dir {
	case DirUp:
		{
			if sm.scrollMenuOffset > 0 {
				sm.scrollMenuOffset--
			}
		}
		return
	case DirDown:
		{
			if sm.scrollMenuOffset < len(sm.buttons)-(sm.height/sm.buttonheight) {
				sm.scrollMenuOffset++
			}
		}
		return
	default:
		return
	}
}
