package gamecommon

import "github.com/hajimehoshi/ebiten/v2"

type ScreenCapture struct {
	filename  string
	mode      CaptureMode
	frame     int
	listening bool
	frames    []*ebiten.Image
}

type CaptureMode int

const (
	GIF CaptureMode = iota
	Picture
)

func NewScreenCapture(frames int, filename string) *ScreenCapture {
	if frames < 1 {
		panic("Can't capture 0 frames")
	}
	if frames < 2 {
		return &ScreenCapture{mode: Picture, filename: filename, frame: 0, frames: make([]*ebiten.Image, 1), listening: false}
	}
	return &ScreenCapture{filename: filename, frame: 0, frames: make([]*ebiten.Image, frames), listening: false}
}

func (sc *ScreenCapture) Capture() {
	sc.listening = true
}

func (sc *ScreenCapture) Draw(screen *ebiten.Image) {
	if sc.listening {
		if sc.frame >= len(sc.frames) {
			sc.listening = false
			sc.frame = 0
			switch sc.mode {
			case Picture:
				FromEbitenFrame(sc.frames[0]).SaveToFile(sc.filename)
			case GIF:
				FromEbitenFrames(sc.frames, sc.filename)
			}
		}
		sc.frames[sc.frame] = ebiten.NewImageFromImage(screen)
		sc.frame++
	}
}
