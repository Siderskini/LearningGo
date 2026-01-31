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

package game

import (
	"home/gamecommon"
	"home/gamecommon/save"
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
			save.SaveGame(g.save)
		}
		nameButton.Text = temp
	}
	return nil
}
