package gamecommon

import "github.com/hajimehoshi/ebiten/v2"

type Animation struct {
	Frames []*ebiten.Image
	frame  int
}

func NewAnimation(frames []*ebiten.Image) *Animation {
	return &Animation{Frames: frames, frame: 0}
}

func (animation *Animation) Update(updateFunc func(int) bool) bool {
	return updateFunc(animation.frame)
}

func (animation *Animation) Draw(screen *ebiten.Image) {
	//Draw animation content here
	if animation.frame >= 1000000 {
		animation.frame = 0
	}
	currentImage := animation.Frames[animation.frame%len(animation.Frames)]
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(currentImage, op)
	animation.frame++
}
