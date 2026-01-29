package gamecommon

import (
	"encoding/gob"
	"os"
)

const (
	fileName = "save.dat"
)

// Takes a pointer to an empty any struct, populates the struct with the save data from the save file, and returns back the poin
func LoadGame(s any) (any, error) {
	return loadGameFromSaveFile(s)
}

func loadGameFromSaveFile(s any) (any, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func loadGameFromBrowserStorage() (any, error) {
	return nil, nil
}

// Takes a pointer to any struct and a filename, and saves the struct data to the given filename
func SaveGame(s any) error {

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&s)
	if err != nil {
		return err
	}
	return nil
}
