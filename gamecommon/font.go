package gamecommon

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Font struct{}

func NewFont(bs []byte) *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
