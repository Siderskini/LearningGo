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
