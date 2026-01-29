package gamecommon

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Audio struct{}

func NewAudio(audioContext *audio.Context, bs []byte) *audio.Player {
	s, err := wav.DecodeF32(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	loop := audio.NewInfiniteLoopF32(s, s.Length())
	audioPlayer, err := audioContext.NewPlayerF32(loop)
	audioPlayer.SetVolume(0.04)
	if err != nil {
		log.Fatal(err)
	}
	return audioPlayer
}
