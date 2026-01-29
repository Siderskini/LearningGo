package save

func LoadGame(s any) (any, error) {
	return LoadGameLowLevel(s)
}

func SaveGame(s any) error {
	return SaveGameLowLevel(s)
}
