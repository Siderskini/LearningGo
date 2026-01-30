package gamecommon

//Basic utility package to transform gifs into ebiten animatables
import (
	"image"
	"image/gif"
	"io/fs"
	"log"
	"os"

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
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = gif.EncodeAll(file, anim)
	if err != nil {
		panic(err)
	}
}
