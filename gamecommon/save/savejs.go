//go:build js

package save

import (
	"encoding/json"
	"os"
	"syscall/js"
)

func LoadGameLowLevel(s any) (any, error) {
	// Get the global window object
	window := js.Global()
	// Access localStorage
	localStorage := window.Get("localStorage")

	// Save a key-value pair
	jssave := localStorage.Call("getItem", "savedat")
	if jssave.IsNull() {
		return nil, os.ErrNotExist
	}
	json.Unmarshal([]byte(jssave.String()), s)
	return s, nil
}

func SaveGameLowLevel(s any) error {
	// Get the global window object
	window := js.Global()
	// Access localStorage
	localStorage := window.Get("localStorage")

	// Save a key-value pair
	jssave, err := json.Marshal(s)
	if err != nil {
		return err
	}
	localStorage.Call("setItem", "savedat", string(jssave))
	return nil
}
