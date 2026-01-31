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

//Basic utility package to transform gifs into ebiten animatables
import (
	"image"
	"image/gif"
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	framerate = 60
)

// ToEbitenFrames loads a gif from the given filename and converts it into a slice of ebiten.Images. Each frame is duplicated to match the given duration in frames.
func ToEbitenFrames(file fs.File, duration int) ([]*ebiten.Image, error) {
	gifImg := TryPanic(gif.DecodeAll(file))
	frameConv := duration/len(gifImg.Image) + 1
	frames := make([]*ebiten.Image, duration)
	for i := 0; i < duration; i++ {
		frames[i] = ebiten.NewImageFromImage(gifImg.Image[i/frameConv])
	}
	return frames, nil
}

// FromEbitenFrames takes a series of ebiten image frames, computes a gif out of them, and dumps it into a file with the filename
func FromEbitenFrames(frames []*ebiten.Image, filename string) {
	gifFrames := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	for _, frame := range frames {
		gifFrames = append(gifFrames, FromEbitenFrame(frame).ToNativeImage())
		delays = append(delays, 0)
	}
	anim := &gif.GIF{
		Image: gifFrames,
		Delay: delays,
	}
	// Write to file
	file := TryLog(os.Create(filename))
	defer file.Close()
	TryPanic("", gif.EncodeAll(file, anim))
}
