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

type Fish struct {
	Item
	Size   int // in centimeters
	Weight int // in grams
}

var fishes = map[string]Fish{
	"RedSnapper": Fish{Item{"Red Snapper", 15}, 30, 2000},
	"Barracuda":  Fish{Item{"Barracuda", 25}, 60, 5000},
	"Clownfish":  Fish{Item{"Clownfish", 10}, 15, 500},
	"Salmon":     Fish{Item{"Salmon", 30}, 70, 7000},
	"Tuna":       Fish{Item{"Tuna", 35}, 80, 8000},
	"Trout":      Fish{Item{"Trout", 20}, 40, 3000},
	"Bass":       Fish{Item{"Bass", 22}, 50, 4000},
	"Carp":       Fish{Item{"Carp", 18}, 45, 3500},
	"Swordfish":  Fish{Item{"Swordfish", 28}, 65, 6000},
}

func GetFish(key string) (Fish, bool) {
	fish, exists := fishes[key]
	return fish, exists
}
