package main

import (
	"errors"
	"fmt"
	"log"
  "time"
  "github.com/sjoerdschouten/adventofcode2020/util"
)

// getProduct checks if an array contains three values
// that add up to 2020 and return their multiplication.
 func getProduct(input []int) (string, error) {   
  var sum int
  for i := 0; i < len(input); i++ {
    for j := i+1; j < len(input); j++ {
      for y := j+1; y < len(input); y++ {
        if (input[i] + input[j] + input[y] ==  2020) {
          sum = input[i] * input[j] * input[y]
          return fmt.Sprint("%d + %d + %d == 2020, multiplied it is: %d", input[i], input[j], input[y], sum), nil
        }
      }
    }     
  }
  return "", errors.New("no combination of elements found that add up to 2020")
}

// getTwentyTwenty checks if an arry contains two values
// that add up to 2020 and returns their multiplication
func getTwentyTwenty(input []int) (string, error) {
  var sum int
  for i := 0; i < len(input); i++ {
    for j := i+1; j < len(input); j++ {
      if (input[i] + input[j] ==  2020) {
        sum = input[i] * input[j]
        return fmt.Sprint("%d + %d == 2020, multiplied it is: %d", input[i], input[j], sum), nil
      }
    }     
  }
  return "", errors.New("no combination of elements found that add up to 2020")
}


func main() {
  // read input
  lines, err := util.ReadInts("./input/input")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }
  // start timer for puzzle 1
  startTime := time.Now()
  // get result
  result, err := getTwentyTwenty(lines)
  if err != nil {
    log.Fatalf("getTwentyTwenty: %s", err)
  }
  // stop timer for puzzle 1
  fmt.Println("Part 1:", result, "Took:", time.Since(startTime))

  startTime = time.Now()
	result, err = getProduct(lines)
  if err != nil { 
    log.Fatal("getProduct: %s", err)
  }
	fmt.Println("Part 2:", result, "Took:", time.Since(startTime))
}
