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

//go:build js

package save

import (
	"encoding/json"
	"home/gamecommon"
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

func SaveGameLowLevel(s any) {
	// Get the global window object
	window := js.Global()
	// Access localStorage
	localStorage := window.Get("localStorage")

	// Save a key-value pair
	jssave := gamecommon.TryPanic(json.Marshal(s))
	localStorage.Call("setItem", "savedat", string(jssave))
}
