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
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames []*ebiten.Image
	frame  int
}

func NewAnimation(resources embed.FS, filename string, duration int) *Animation {
	file := TryPanic(resources.Open(filename))
	Frames := TryPanic(ToEbitenFrames(file, duration))
	return &Animation{Frames: Frames, frame: 0}
}

func (animation *Animation) Update(updateFunc func(int) bool) bool {
	if animation.frame >= 1000000 {
		animation.frame = 0
	}
	animation.frame++
	ret := updateFunc(animation.frame)
	return ret
}

func (animation *Animation) Draw(screen *ebiten.Image) {
	//Draw animation content here
	currentImage := animation.Frames[animation.frame%len(animation.Frames)]
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(currentImage, op)
}
