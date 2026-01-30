package gamecommon

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
