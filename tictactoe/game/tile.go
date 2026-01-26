// Copyright 2016 The Ebiten Authors
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
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

// TileData represents a tile information like a value and a position.
type TileData struct {
	value string
	x     int
	y     int
}

// Tile represents a tile information including TileData and animation states.
type Tile struct {
	current    TileData
	hasBeenSet bool
}

// Pos returns the tile's current position.
// Pos is used only at testing so far.
func (t *Tile) Pos() (int, int) {
	return t.current.x, t.current.y
}

// Value returns the tile's current value.
// Value is used only at testing so far.
func (t *Tile) Value() string {
	return t.current.value
}

func (t *Tile) Update(value string) {
	t.current.value = value
	t.hasBeenSet = true
}

// NewTile creates a new Tile object.
func NewTile(value string, x, y int) *Tile {
	return &Tile{
		current: TileData{
			value: value,
			x:     x,
			y:     y,
		},
		hasBeenSet: false,
	}
}

const (
	tileSize   = 80
	tileMargin = 4
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}

// Draw draws the current tile to the given boardImage.
func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.current.x, t.current.y
	v := t.current.value
	if v == "" {
		return
	}
	op := &ebiten.DrawImageOptions{}
	x := i*tileSize + (i+1)*tileMargin
	y := j*tileSize + (j+1)*tileMargin

	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(tileBackgroundColor(v))
	boardImage.DrawImage(tileImage, op)

	size := 48.0

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(x)+float64(tileSize)/2, float64(y)+float64(tileSize)/2)
	textOp.ColorScale.ScaleWithColor(tileColor(v))
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	text.Draw(boardImage, v, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   size,
	}, textOp)
}
