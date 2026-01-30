package gamecommon

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Font struct{}

func NewFont(bs []byte) *text.GoTextFaceSource {
	return TryPanic(text.NewGoTextFaceSource(bytes.NewReader(bs)))
}
