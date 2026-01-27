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

func (button *Button) Draw(screen *ebiten.Image) {
	x := button.shape.X
	y := button.shape.Y

	size := 12.0
	textOp := &text.DrawOptions{}
	textOp.ColorScale.ScaleWithColor(button.Color)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter

	switch g := button.shape.geometry.(type) {
	case Rectangle:
		textOp.GeoM.Translate(float64(x+g.Width/2), float64(y+g.Height/2))
		vector.FillRect(screen, float32(x), float32(y), float32(g.Width), float32(g.Height), button.BackGroundColor, true)
	case Circle:
		vector.FillCircle(screen, float32(x), float32(y), float32(g.Radius), button.BackGroundColor, true)
	}

	text.Draw(screen, button.Text, &text.GoTextFace{
		Source: button.faceSource,
		Size:   size,
	}, textOp)
}
