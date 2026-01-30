//go:build !js

package save

import (
	"encoding/gob"
	"home/gamecommon"
	"os"
)

const (
	fileName = "save.dat"
)

// Takes a pointer to an empty any struct, populates the struct with the save data from the save file, and returns back the poin
func LoadGameLowLevel(s any) (any, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	gamecommon.TryPanic("", decoder.Decode(&s))
	return s, nil
}

// Takes a pointer to any struct and a filename, and saves the struct data to the given filename
func SaveGameLowLevel(s any) {
	file := gamecommon.TryPanic(os.Create(fileName))
	defer file.Close()
	encoder := gob.NewEncoder(file)
	gamecommon.TryPanic("", encoder.Encode(&s))
}
