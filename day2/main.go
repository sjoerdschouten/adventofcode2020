package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	util "github.com/sjoerdschouten/adventofcode2020/utils"
)

func main() {
  input, err := util.ReadLines("./input/input.txt")
  if err != nil {
    log.Fatalf("ReadLines(): %s", err)
  }

  startTime := time.Now()
  result := solvePuzzleOne(input)
  fmt.Println("Part 1:", result, "Took:", time.Since(startTime))
  
  result = solvePuzzleTwo(input)
  startTime = time.Now()
  fmt.Println("Part 2:", result, "Took:", time.Since(startTime))
}

type Policy struct {
  Minimum int
  Maximum int
}

type Set struct {
  Policy
  Letter string
  Password string
}


func solvePuzzleOne(input []string) string{
  count := 0
  for _, line := range input {
    set := parseLineToSet(line)
    if (set.satisfiesPolicy()) {
      count++
    }
  }
  return fmt.Sprint(count)
}

func solvePuzzleTwo(input []string) string {
  count := 0
  for _, line := range input {
    set := parseLineToSet(line)
    if (set.satisfiesEnhancedPolicy()) {
      count++
    }
  }
  return fmt.Sprint(count)
}


// parseSet return an array with four elements representing the low, high, current letter and passsword
func parseLineToSet(set string) *Set {
  split := strings.Fields(set)
  minMax :=  strings.Split(split[0], "-")
  min, _ := strconv.Atoi(minMax[0])
  max, _ := strconv.Atoi(minMax[1])
  
  letter := string(split[1][0])

  passWord := split[2] 
  
  return &Set{
    Policy: Policy{
      Minimum: min,
      Maximum: max, 
    },
    Letter: letter,
    Password: passWord,
  }
}

func (set *Set)satisfiesPolicy() bool {
  letterCount := strings.Count(set.Password, set.Letter)
  return letterCount >= set.Minimum && letterCount <= set.Maximum
}

func (set *Set)satisfiesEnhancedPolicy() bool {
  if (string(set.Password[set.Policy.Minimum -1]) == set.Letter && string(set.Password[set.Maximum -1]) != set.Letter) {
    return true
  } else if (string(set.Password[set.Minimum -1]) != set.Letter && string(set.Password[set.Maximum -1]) == set.Letter) {
    return true
  } else {
    return false
  }
}
