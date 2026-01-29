# Generic Fishing Game
Making a generic fishing game with ebitengine. Code contains excerpts edited from https://github.com/hajimehoshi/ebiten/tree/main/examples/.
The purpose of this project is to create the generic building blocks for a game:
- Title screen
- Menu with buttons
- Save functionality
- Animations
- Audio

There is a possibility of expanding this scope in the future. As buiding blocks are made, they will be added to gamecommon for use in future projects.

UPDATE: Above scope is complete in a basic form, but can be improved.

Remaining work includes:
- Support for different platforms (windows, macos, browser, mobile)
- Icons, textures, more content
- Support for more functionality (textanimations, complex animations, gif production, etc)

Build:

Linux (supported):  `GOOS=linux GOARCH=amd64 go build -o fishingapp main.go`

Windows (supported):  `GOOS=linux GOARCH=amd64 go build -o fishing.exe main.go`

Browser (supported):  `GOOS=js GOARCH=wasm go build -o fishing/web/main.wasm main.go`

MacOS (supported):  `GOOS=darwin GOARCH=arm64 go build -o fishingapp main.go` (I was not able to build this on linux, only MacOS)

More on the browser deployment in the web folder.