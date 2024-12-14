package main

import (
	"fmt"
	"strings"

	"valyntyler.com/aoc-2024-day-14/input"
	. "valyntyler.com/aoc-2024-day-14/src/robot"
)

func main() {
  fmt.Println("Solving part 1...")

  // var robots []Robot

  var s string = string(input.Bytes)
  var result string = strings.Trim(s, " \n")
  var lines []string = strings.Split(result, "\n")

  for _, line := range lines {
    robot := ParseRobot(line)
    fmt.Println(robot)
  }
}
