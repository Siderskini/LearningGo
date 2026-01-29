package game

import (
	"embed"
	_ "embed"
)

var (
	//go:embed fishing.gif fishing.wav
	resources embed.FS
)
