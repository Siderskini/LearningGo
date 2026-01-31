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
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	Name            string
	shape           Shape
	input           *Input
	Text            string
	faceSource      *text.GoTextFaceSource
	Color           color.Color
	BackGroundColor color.Color
	Icon            *ebiten.Image
}

func NewRectangleButton(name string, x, y, width, height int, text string, faceSource *text.GoTextFaceSource, color color.Color, backgroundColor color.Color, input *Input) Button {
	return Button{
		Name:            name,
		shape:           NewRectangle(x, y, width, height),
		Text:            text,
		faceSource:      faceSource,
		input:           input,
		Color:           color,
		BackGroundColor: backgroundColor,
	}
}

func NewCircleButton(name string, x, y, radius int, text string, faceSource *text.GoTextFaceSource, color color.Color, backgroundColor color.Color, input *Input) Button {
	return Button{
		Name:            name,
		shape:           NewCircle(x, y, radius),
		Text:            text,
		faceSource:      faceSource,
		input:           input,
		Color:           color,
		BackGroundColor: backgroundColor,
	}
}

func (button *Button) IsPressed() bool {
	pressedLocation, ok := button.input.PressedLocation()
	if !ok {
		return false
	}
	return button.shape.Contains(pressedLocation[0], pressedLocation[1])
}

// If a button is being drawn onto an image other than the screen, you can pass the x and y of that image here to handle the translation
func (button *Button) IsPressedBounded(x, y, width, height int) bool {
	pressedLocation, ok := button.input.PressedLocation()
	if !ok {
		return false
	}
	if !NewRectangle(x, y, width, height).Contains(pressedLocation[0], pressedLocation[1]) {
		return false
	}
	return button.shape.Contains(pressedLocation[0]-x, pressedLocation[1]-y)
}

func (button *Button) Draw(screen *ebiten.Image) {
	x := button.shape.X
	y := button.shape.Y

	size := 12.0
	textOp := &text.DrawOptions{}
	textOp.ColorScale.ScaleWithColor(button.Color)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	switch g := button.shape.Geometry.(type) {
	case Rectangle:
		textOp.GeoM.Translate(float64(x+g.Width/2), float64(y+g.Height/2))
		vector.FillRect(screen, float32(x), float32(y), float32(g.Width), float32(g.Height), button.BackGroundColor, true)
	case Circle:
		vector.FillCircle(screen, float32(x), float32(y), float32(g.Radius), button.BackGroundColor, true)
	}

	if button.Icon != nil {
		iconOp := &ebiten.DrawImageOptions{}
		iconOp.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(button.Icon, iconOp)
	}

	text.Draw(screen, button.Text, &text.GoTextFace{
		Source: button.faceSource,
		Size:   size,
	}, textOp)
}
