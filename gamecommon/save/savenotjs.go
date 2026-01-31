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
