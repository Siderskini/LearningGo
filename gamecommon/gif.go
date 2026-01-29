package gamecommon

//Basic utility package to transform gifs into ebiten animatables
import (
	"image/gif"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	framerate = 60
)

// ToEbitenFrames loads a gif from the given filename and converts it into a slice of ebiten.Images. Each frame is duplicated to match the given duration in frames.
func ToEbitenFrames(file fs.File, duration int) ([]*ebiten.Image, error) {

	gifImg, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}
	frameConv := duration/len(gifImg.Image) + 1
	frames := make([]*ebiten.Image, duration)
	for i := 0; i < duration; i++ {
		frames[i] = ebiten.NewImageFromImage(gifImg.Image[i/frameConv])
	}
	return frames, nil
}
