package main

import (
	"log"
	"time"
  "fmt"

	util "github.com/sjoerdschouten/adventofcode2020/utils"
)

func main() {
  input, err := util.ReadLines("input/input.txt")
  if err != nil {
    log.Fatalf("ReadLines(): %s", err)
  }
  startTime := time.Now()
  result := solvePuzzleOne(input)
  fmt.Println("Part 1:", result, "Took:", time.Since(startTime))
}

func solvePuzzleOne(input []string) (answer int) {
    rightAdd := 3
    downAdd := 1
    currentRight := rightAdd
    currentDown := downAdd
    rightMax := len(input[0])
    tree := '#'
    
    for currentDown < len(input) {
      if input[currentDown][currentRight] == byte(tree) {
        answer ++
      }

      currentDown = currentDown + downAdd
      currentRight = (currentRight + rightAdd) % rightMax
    }
    return
}
