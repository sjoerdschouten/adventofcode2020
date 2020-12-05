package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	util "github.com/sjoerdschouten/adventofcode2020/utils"
)

func main() {
  startTime := time.Now()
  input, err := util.ReadLines("./input/input.txt")
  if err != nil {
    log.Fatalf("ReadLines(): %s", err)
  }
  fmt.Println("Part 1:", solvePuzzleOne(input), "Took:", time.Since(startTime))
  
  startTime = time.Now()
  fmt.Println("Part 2:", solvePuzzleTwo(input), "Took:", time.Since(startTime))
}

func solvePuzzleOne(input []string) int {
  var valid int
  for _, line := range input {
    set := parseLineToSet(line)
    if (satisfiesPolicy(set)) {
      valid++
    }
  }
  return valid
}

func solvePuzzleTwo(input []string) int {
  var valid int
  for _, line := range input {
    set := parseLineToSet(line)
    if (satisfiesEnhancedPolicy(set)) {
      valid++
    }
  }
  return valid
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

// parseSet return an array with four elements representing the low, high, current letter and passsword
func parseLineToSet(set string) *Set {
  re := regexp.MustCompile(`^\d+`)
  minimum, _ := strconv.Atoi(re.FindString(set))

  re = regexp.MustCompile(`\d+( )`)
  highString := re.FindString(set)
  highString = strings.Trim(highString, " ")
  maximum, _ := strconv.Atoi(highString)

  re = regexp.MustCompile(`\w(:)`)
  letter := re.FindString(set)[:1]

  re = regexp.MustCompile(`[a-z]\w+`)
  passWord := re.FindString(set)
  
  setStruct := &Set{
    Policy: Policy{
      Minimum: minimum,
      Maximum: maximum,
    },
    Letter: letter,
    Password: passWord,
  }
  return setStruct
}

func satisfiesPolicy(set *Set) bool {
  letterCount := strings.Count(set.Password, set.Letter)
  return letterCount >= set.Minimum && letterCount <= set.Maximum
}

func satisfiesEnhancedPolicy(set *Set) bool {
  if (string(set.Password[set.Minimum -1]) == set.Letter && string(set.Password[set.Maximum -1]) != set.Letter) {
    return true
  } else if (string(set.Password[set.Minimum -1]) != set.Letter && string(set.Password[set.Maximum -1]) == set.Letter) {
    return true
  } else {
    return false
  }
}
