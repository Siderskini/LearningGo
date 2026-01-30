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
