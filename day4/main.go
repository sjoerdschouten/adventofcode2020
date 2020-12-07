package main

import (
	"fmt"
	"log"

	util "github.com/sjoerdschouten/adventofcode2020/utils"
)



func main() {
  input, err := util.ReadNonEmptyLines("input/test.txt")
  if err != nil {
    log.Fatalln(err)
  }
  for _, line := range input {
    fmt.Println(line)
  }
}
