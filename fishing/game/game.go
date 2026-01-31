// Copyright 2026 Siddharth Viswnathan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package game

import (
	"encoding/gob"
	"home/gamecommon"
	"home/gamecommon/save"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Global constants
const (
	ScreenWidth   = 420
	ScreenHeight  = 600
	fontSize      = 24
	smallFontSize = fontSize / 2
)

// Game represents the game state
type Game struct {
	mode Mode
	save *Save
}

/*
Save data contains:
Name: the player name
Fish: the currently held fish. the keys will include all caught fish
Inventory: the currently held non-fish items
Money: the player's balance
*/
type Save struct {
	Name      string
	Fish      map[string]int
	Inventory map[string]int
	Money     int
}

// Current mode of the game
type Mode int

const (
	Title Mode = iota
	Animation
	Shopping
	Fishing
	Initializing
)

// Resource Variables
var (
	arcadeFaceSource *text.GoTextFaceSource
	input            *gamecommon.Input
	audioContext     *audio.Context
	loop             *audio.InfiniteLoop
	audioPlayer      *audio.Player
	screenCapture    *gamecommon.ScreenCapture
)

func init() {
	gob.Register(&Save{})
	input = gamecommon.NewInput()
	arcadeFaceSource = gamecommon.NewFont(fonts.MPlus1pRegular_ttf)
	backGroundAnimation = gamecommon.NewAnimation(resources, "background.gif", 640)
	fishingAnimation = gamecommon.NewAnimation(resources, "fishing.gif", 120)
	audioContext = audio.NewContext(48000)
	audioPlayer = gamecommon.NewAudio(audioContext, fishingwav)
	screenCapture = gamecommon.NewScreenCapture(10, "temp.gif")
}

// Game variables
var (
	shop                *Shop
	titlePage           *TitlePage
	activity            *Activity
	initial             *Initial
	backGroundAnimation *gamecommon.Animation
	fishingAnimation    *gamecommon.Animation
)

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	m := Title
	// Try to load a save. If one doesn't exist, send the user to initializing
	save, err := save.LoadGame(&Save{})
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
	g := &Game{
		save: save.(*Save),
		mode: m,
	}
	shop = NewShop()
	titlePage = &TitlePage{}
	activity = &Activity{}
	initial = &Initial{}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func fishingAnimationUpdate(i int) bool {
	return i > 120
}

// Update updates the current game state.
func (g *Game) Update() error {
	input.Update()
	audioPlayer.Play()
	switch g.mode {
	case Title:
		backGroundAnimation.Update(func(i int) bool { return true })
		return titlePage.Update(g)
	case Animation:
		if fishingAnimation.Update(fishingAnimationUpdate) {
			screenCapture.Capture()
			g.mode = Title
		}
		return nil
	case Shopping:
		backGroundAnimation.Update(func(i int) bool { return true })
		return shop.Update(g)
	case Fishing:
		backGroundAnimation.Update(func(i int) bool { return true })
		return activity.Update(g)
	case Initializing:
		backGroundAnimation.Update(func(i int) bool { return true })
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
		backGroundAnimation.Draw((screen))
		initial.Draw(screen)
	}
	screenCapture.Draw(screen)
}
