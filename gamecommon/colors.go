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

package gamecommon

import (
	"image/color"
)

var (
	BackgroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
)

func tileColor(value string) color.Color {
	switch value {
	case "O":
		return color.RGBA{0x00, 0x00, 0xee, 0x59}
	case "X":
		return color.RGBA{0xee, 0x00, 0x00, 0x59}
	}
	panic("not reach")
}

func tileBackgroundColor(value string) color.Color {
	return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
}
