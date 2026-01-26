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
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var taskTerminated = errors.New("tictactoe: task terminated")

type task func() error

// Board represents the game board.
type Board struct {
	size   int
	tiles  map[int]*Tile
	tasks  []task
	player string
}

// NewBoard generates a new Board with giving a size.
func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:   size,
		tiles:  make(map[int]*Tile),
		player: "X",
	}
	for i := 0; i < size*size; i++ {
		b.tiles[i] = NewTile("", i%size, i/size)
	}
	return b, nil
}

// Update updates the board state.
func (b *Board) Update(input *Input) (error, bool) {
	if 0 < len(b.tasks) {
		t := b.tasks[0]
		if err := t(); err == taskTerminated {
			b.tasks = b.tasks[1:]
		} else if err != nil {
			return err, false
		}
		return nil, false
	}
	if pressed, ok := input.PressedTile(); ok {
		if err, won := b.PressTile(pressed); err != nil {
			return err, false
		} else if won {
			return nil, true
		}
	}
	return nil, false
}

// Converts a mouse position to a tile position using cursor x and y as input and returning tile x and y as output
func mousePositionToTilePosition(b *Board, x int, y int) int {
	// If the position is in top of left margin, it's not a tile
	tileY := x/(tileSize+tileMargin) - 1
	tileX := y/(tileSize+tileMargin) - 2
	ret := tileX*b.size + tileY
	if ret < 0 || ret >= b.size*b.size {
		return -1
	}
	return tileX*b.size + tileY

}

func checkVictory(b *Board) bool {
	for i := 0; i < b.size; i++ {
		// Check rows
		if b.tiles[i*b.size].Value() != "" && b.tiles[i*b.size].Value() == b.tiles[i*b.size+1].Value() && b.tiles[i*b.size].Value() == b.tiles[i*b.size+2].Value() {
			return true
		}
		//Check the columns
		if b.tiles[i].Value() != "" && b.tiles[i].Value() == b.tiles[i+3].Value() && b.tiles[i].Value() == b.tiles[i+6].Value() {
			return true
		}
	}
	// Check diagonals
	if b.tiles[0].Value() != "" && b.tiles[0].Value() == b.tiles[4].Value() && b.tiles[0].Value() == b.tiles[8].Value() {
		return true
	}
	if b.tiles[2].Value() != "" && b.tiles[2].Value() == b.tiles[4].Value() && b.tiles[2].Value() == b.tiles[6].Value() {
		return true
	}
	return false
}

func (b *Board) PressTile(pos []int) (error, bool) {
	position := mousePositionToTilePosition(b, pos[0], pos[1])
	if position < 0 {
		return nil, false
	}
	if !b.tiles[position].hasBeenSet {
		b.tiles[position].Update(b.player)
		if checkVictory(b) {
			return nil, true
		}
		if b.player == "X" {
			b.player = "O"
		} else {
			b.player = "X"
		}
	}
	return nil, false
}

// Size returns the board size.
func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)
	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			v := ""
			op := &ebiten.DrawImageOptions{}
			x := i*tileSize + (i+1)*tileMargin
			y := j*tileSize + (j+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			op.ColorScale.ScaleWithColor(tileBackgroundColor(v))
			boardImage.DrawImage(tileImage, op)
		}
	}
	for i := 0; i < b.size*b.size; i++ {
		b.tiles[i].Draw(boardImage)
	}
}
