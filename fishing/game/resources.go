package game

import (
	"embed"
	_ "embed"
)

var (
	//go:embed fishing.gif background.gif
	resources embed.FS

	//go:embed fishing.wav
	fishingwav []byte
)
