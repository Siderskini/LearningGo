package gamecommon

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageResource struct {
	Image *ebiten.Image
}

var palette = []color.Color{
	color.RGBA{R: 255, G: 0, B: 0, A: 255},     // Red
	color.RGBA{R: 0, G: 255, B: 0, A: 255},     // Green
	color.RGBA{R: 0, G: 0, B: 255, A: 255},     // Blue
	color.RGBA{R: 255, G: 255, B: 0, A: 255},   // Yellow
	color.RGBA{R: 255, G: 255, B: 255, A: 255}, // White
	color.RGBA{R: 0, G: 0, B: 0, A: 255},       // Black
}

func NewImage(bs []byte) *ImageResource {
	img, _, err := image.Decode(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	return &ImageResource{Image: ebiten.NewImageFromImage(img)}
}

func FromEbitenFrame(i *ebiten.Image) *ImageResource {
	return &ImageResource{Image: i}
}

func (i *ImageResource) ToNativeImage() *image.Paletted {
	// Create a new image.Paletted from the pixel data
	paletted := image.NewPaletted(i.Image.Bounds(), palette)
	draw.Draw(paletted, i.Image.Bounds(), i.Image, i.Image.Bounds().Min, draw.Src)
	return paletted
}

func (i *ImageResource) SaveToFile(filename string) {
	// Write to file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, i.ToNativeImage())
	if err != nil {
		panic(err)
	}
}
