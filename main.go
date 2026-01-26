package main

import (
	"home/fishing"
	"home/hello"
	"home/tictactoe"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hello":
		hello.Hello()
		break
	case "tictactoe":
		tictactoe.StartGame()
		break
	case "fishing":
		fishing.StartGame()
		break
	default:
		panic("Invalid Argument")
	}
}
