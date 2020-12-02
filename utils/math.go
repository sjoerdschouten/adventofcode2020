package util

import (
  "strconv"
)

// GetInt returns the given string as an int, or panics if it is invalid
func GetInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return res
}
