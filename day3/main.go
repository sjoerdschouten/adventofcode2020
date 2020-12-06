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
  result := solve1(input)
  fmt.Println("Part 1:", result, "Took:", time.Since(startTime))
  startTime = time.Now()
  result = solve2(input)
  fmt.Println("Part 2:", result, "Took:", time.Since(startTime))
}

type Slope struct {
  Right int
  Down int
}

// Returns the amount of trees hit given a slope
func (slope *Slope) treesHit(input []string) (hit int) {
  x := slope.Right
  y := slope.Down
  for y < len(input) {
    if input[y][x] == byte('#') {
      hit++
    }
    x = (x + slope.Right) % len(input[0])
    y += slope.Down
  }
  return
}

func solve1(input []string) (answer int) {
  slope := &Slope {
    Right: 3,
    Down: 1,
  }
  answer = slope.treesHit(input)
  return
}

func solve2(input []string) (answer int) {
  slopes := []*Slope {
    {1,1},
    {3,1},
    {5,1},
    {7,1},
    {1,2},
  }
  answer = 1
  for _, slope := range slopes {
    answer *= slope.treesHit(input)
  }
  return
}
