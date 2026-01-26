package game

type Fish struct {
	Name   string
	Size   int // in centimeters
	Weight int // in grams
	Price  int
}

var fishes = map[string]Fish{
	"RedSnapper": Fish{"Red Snapper", 30, 2000, 15},
	"Barracuda":  Fish{"Barracuda", 60, 5000, 25},
	"Clownfish":  Fish{"Clownfish", 15, 500, 10},
	"Salmon":     Fish{"Salmon", 70, 7000, 30},
	"Tuna":       Fish{"Tuna", 80, 8000, 35},
	"Trout":      Fish{"Trout", 40, 3000, 20},
	"Bass":       Fish{"Bass", 50, 4000, 22},
	"Carp":       Fish{"Carp", 45, 3500, 18},
}

func GetFish(key string) (Fish, bool) {
	fish, exists := fishes[key]
	return fish, exists
}
