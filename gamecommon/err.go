package gamecommon

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

import "log"

type Err struct{}

// So many functions return (val, error), only for us to log or ignore the error
// This function helps reduce this by allowing for a lambda to take care of any error handling
func TryCatch[T any](catch func(T, error) T) func(T, error) T {
	return func(val T, err error) T {
		if err != nil {
			return catch(val, err)
		}
		return val
	}
}

// This is functionally the result of TryCatch(log.Fatal+return nil)
func TryLog[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err)
		var zero T
		return zero
	}
	return val
}

// This is functionally the same as TryCatch(panic)
func TryPanic[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
