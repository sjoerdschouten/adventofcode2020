package main

import (
	"fmt"
	"strings"
	"time"

	util "github.com/sjoerdschouten/adventofcode2020/utils"
	"golang.org/x/text/cases"
)

func main() {
  input := util.ReadEntireFile("input/input.txt")
  passports := parseLines(input)
  startTime := time.Now()
  result := solvePuzzle1(passports)
  fmt.Println("Part 1:", result, "Took:", time.Since(startTime))
}

func parseLines(input string) []map[string]string {
  passports := make([]map[string]string, 0)
  for _, values := range strings.Split(input, "\n\n") {
    // for each passport line we create a unique map that we can append to the result.
    passport := make(map[string]string)
    for _, value := range strings.Fields(values) { 
      // for each passport value we split and add the key and value to the map.
      // if there are no more value we append the passport map to the slice that contains a map of passports.
      kv := strings.Split(value, ":") 
      passport[kv[0]] = kv[1]
		}
      passports = append(passports, passport)
  }
  fmt.Println(len(passports))

  return passports
}

func solvePuzzle1(input []map[string]string) (result int) {
  for _, passport := range input {
    _, hgt := passport["hgt"]
    _, byr := passport["byr"]
    _, cid := passport["cid"]
    if !cid && len(passport) == 7 {

      result ++
      continue
    }
    if hgt && byr && len(passport) == 8 {
      fmt.Println(passport)
      result++
    }
  } 
  return
}

func solvePuzzle2(input []map[string]string) (result int) {
  for _, passport := range input {
    _, hgt := passport["hgt"]
    _, byr := passport["byr"]
    _, cid := passport["cid"]
    if !cid && len(passport) == 7 {
      valid := true
      if hgt < 1920 || hgt > 2002 {
        valid = false
      }
      if v, :+ := m["pi"]; found {
       fmt.Println(v)
      
      }

      result ++
      continue
    }
    if hgt && byr && len(passport) == 8 {
      fmt.Println(passport)
      result++
    }
  } 
  return
}



// mandatory keys
// hgt
// byr


// optional cid


// mandatory keys
// hgt
// byr


// optional cid
