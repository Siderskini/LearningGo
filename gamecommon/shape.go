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

type Shape struct {
	X, Y     int
	Geometry any
}

type Rectangle struct {
	Width, Height int
}

type Circle struct {
	Radius int
}

// x and y represent the top-left corner for Rectangle
func NewRectangle(x, y, width, height int) Shape {
	return Shape{
		X: x,
		Y: y,
		Geometry: Rectangle{
			Width:  width,
			Height: height,
		},
	}
}

// x and y represent the center for Circle
func NewCircle(x, y, radius int) Shape {
	return Shape{
		X: x,
		Y: y,
		Geometry: Circle{
			Radius: radius,
		},
	}
}

func (s Shape) Contains(px, py int) bool {
	switch g := s.Geometry.(type) {
	case Rectangle:
		return px >= s.X && px <= s.X+g.Width && py >= s.Y && py <= s.Y+g.Height
	case Circle:
		dx := px - s.X
		dy := py - s.Y
		return dx*dx+dy*dy <= g.Radius*g.Radius
	default:
		return false
	}
}
