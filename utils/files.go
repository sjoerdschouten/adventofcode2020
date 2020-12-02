package util

import (
	"bufio"
	"fmt"
	"os"
)

// Readlines read all lines from a provide filepath and returns a slice of strings
// if the filepath does not exist it will return an error.
func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, fmt.Errorf("ReadLines(): %s", err)
  }    
  defer file.Close()
    
  scanner := bufio.NewScanner(file)
  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, nil
}

// ReadIntegers reads a file from a provided filepath and returns a slice of ReadIntegers.
// Return an error when unable to read a file.
func ReadIntegers(name string) ([]int, error) {
  lines, err := ReadLines(name)
  if err != nil {
    return nil, fmt.Errorf("ReadIntegers(): %s", err)
  }
  var out []int
  for _, line := range(lines) {
    out = append(out, GetInt(line))
  }
  return out, nil
}
