package util

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "io/ioutil"

)

//
func ReadEntireFile(name string) string {
  result, err := ioutil.ReadFile(name)
  if err != nil {
    panic(err)
  }

  return string(result)
}

// ReadNonEmptyLines reada all lines from a provided filepath and returns a slice of string
// an empty line in the source file is used to seperate each element.
func ReadNonEmptyLines(path string) ([][]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, fmt.Errorf("ReadLines(): %s", err)
  }    
  defer file.Close()
  buffer := [][]string{}
  block := []string{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    if len(strings.TrimSpace(line)) != 0 {
        block = append(block, line)
        continue
    }

    // At this point, the script has reached an empty line,
    // which means the block is ready to be processed.
    // If the block is not empty, append it to the buffer and empty it.
    if len(block) != 0 {
        buffer = append(buffer, block)
        block = []string{}
    }
  }
  if len(block) != 0 {
    buffer = append(buffer, block)
  }
  return buffer, nil
}

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
