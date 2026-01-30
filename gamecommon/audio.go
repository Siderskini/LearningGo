package gamecommon

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Audio struct{}

func NewAudio(audioContext *audio.Context, bs []byte) *audio.Player {
	s := TryPanic(wav.DecodeF32(bytes.NewReader(bs)))
	loop := audio.NewInfiniteLoopF32(s, s.Length())
	audioPlayer := TryPanic(audioContext.NewPlayerF32(loop))
	audioPlayer.SetVolume(0.04)
	return audioPlayer
}
