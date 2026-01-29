package game

import (
	"bytes"
	"encoding/gob"
	"home/gamecommon"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth   = 420
	ScreenHeight  = 600
	boardSize     = 3
	titleFontSize = fontSize * 1.25
	fontSize      = 24
	smallFontSize = fontSize / 2
)

// Game represents a game state.
type Game struct {
	mode       Mode
	input      *gamecommon.Input
	boardImage *ebiten.Image
	save       *Save
}

type Mode int

const (
	Title Mode = iota
	Animation
	Shopping
	Fishing
	Initializing
)

func init() {
	gob.Register(save)
	input = gamecommon.NewInput()
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFaceSource = s
	file, err := resources.Open("background.gif")
	backgroundAnimationFrames, err = gamecommon.ToEbitenFrames(file, 640)
	if err != nil {
		log.Fatal(err)
	}
	backGroundAnimation = gamecommon.NewAnimation(backgroundAnimationFrames)

	file, err = resources.Open("fishing.gif")
	fishingAnimationFrames, err = gamecommon.ToEbitenFrames(file, 120)
	if err != nil {
		log.Fatal(err)
	}
	fishingAnimation = gamecommon.NewAnimation(fishingAnimationFrames)

	if audioContext == nil {
		audioContext = audio.NewContext(48000)
	}
	bs, err := resources.ReadFile("fishing.wav")
	if err != nil {
		log.Fatal(err)
	}
	jabD, err := wav.DecodeF32(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	loop = audio.NewInfiniteLoopF32(jabD, jabD.Length())
	audioPlayer, err = audioContext.NewPlayerF32(loop)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	arcadeFaceSource          *text.GoTextFaceSource
	fishingAnimationFrames    []*ebiten.Image
	backgroundAnimationFrames []*ebiten.Image
	framecounter              int
	input                     *gamecommon.Input
	audioContext              *audio.Context
	loop                      *audio.InfiniteLoop
	audioPlayer               *audio.Player
)

type Save struct {
	Name      string
	Fish      map[string]int
	Inventory map[string]int
	Money     int
}

var save *Save
var shop *Shop
var titlePage *TitlePage
var activity *Activity
var initial *Initial
var backGroundAnimation *gamecommon.Animation
var fishingAnimation *gamecommon.Animation

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	m := Title
	save, err := gamecommon.LoadGame(save)
	if err != nil {
		if os.IsNotExist(err) {
			save = &Save{
				Name:      "",
				Fish:      make(map[string]int),
				Inventory: make(map[string]int),
				Money:     0,
			}
			m = Initializing
		} else {
			panic(err)
		}
	}

	shop = &Shop{
		selectedFish:      "",
		selectedItem:      "",
		quantity:          0,
		displayedQuantity: "0",
	}

	titlePage = &TitlePage{}
	activity = &Activity{}
	initial = &Initial{}
	g := &Game{
		input: input,
		save:  save.(*Save),
		mode:  m,
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func fishingAnimationUpdate(i int) bool {
	return i > len(fishingAnimationFrames)
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.input.Update()
	audioPlayer.Play()
	switch g.mode {
	case Title:
		return titlePage.Update(g)
	case Animation:
		if fishingAnimation.Update(fishingAnimationUpdate) {
			g.mode = Title
		}
		return nil
	case Shopping:
		return shop.Update(g)
	case Fishing:
		return activity.Update(g)
	case Initializing:
		return initial.Update(g)
	}
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case Title:
		backGroundAnimation.Draw(screen)
		titlePage.Draw(screen)
	case Animation:
		fishingAnimation.Draw(screen)
	case Shopping:
		backGroundAnimation.Draw(screen)
		shop.Draw(g, screen)
	case Fishing:
		backGroundAnimation.Draw(screen)
		activity.Draw(screen)
	case Initializing:
		initial.Draw(screen)
	}
}
