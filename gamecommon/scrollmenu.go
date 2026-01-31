// Copyright 2026 Siddharth Viswnathan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gamecommon

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScrollMenu struct {
	x, y, width, height, buttonheight int
	buttons                           []*Button
	input                             *Input
	scrollMenuOffset                  int
	selected                          string
}

/*
Create a new ScrollMenu with the given buttons and dimensions.
buttons: slice of Button objects to display in the menu
*/
func NewScrollMenu(buttons []*Button, title string, x, y, width, height, buttonheight int, input *Input) *ScrollMenu {
	for i, button := range buttons {
		switch g := button.shape.Geometry.(type) {
		case Rectangle:
			g.Width = width - 20
			g.Height = buttonheight - 2
			button.shape.Geometry = g
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
		selected:         "",
	}
}

func (sm *ScrollMenu) Draw(screen *ebiten.Image) {
	menuImage := ebiten.NewImage(sm.width, sm.height)
	menuImage.Fill(color.Black)
	for i, button := range sm.buttons {
		button.shape.Y = (i-sm.scrollMenuOffset)*sm.buttonheight + 10
		if sm.selected == button.Name {
			button.Color = color.RGBA{255, 255, 0, 255}
		} else {
			button.Color = color.White
		}
		button.Draw(menuImage)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sm.x), float64(sm.y))
	screen.DrawImage(menuImage, op)
}

func (sm *ScrollMenu) HandleInput() string {
	for _, button := range sm.buttons {
		pressed := button.IsPressedBounded(sm.x, sm.y, sm.width, sm.height)
		if pressed {
			sm.selected = button.Name
			return sm.selected
		}
	}

	//Handle the scrolling
	dir, pressed := sm.input.Dir()
	if !pressed {
		return sm.selected
	}
	switch dir {
	case DirUp:
		{
			if sm.scrollMenuOffset > 0 {
				sm.scrollMenuOffset--
			}
		}
		return sm.selected
	case DirDown:
		{
			if sm.scrollMenuOffset < len(sm.buttons)-(sm.height/sm.buttonheight) {
				sm.scrollMenuOffset++
			}
		}
		return sm.selected
	default:
		return sm.selected
	}
}
