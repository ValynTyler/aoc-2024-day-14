package main

import (
	"fmt"
	"strings"

	"valyntyler.com/aoc-2024-day-14/input"
	. "valyntyler.com/aoc-2024-day-14/src/robot"
	. "valyntyler.com/aoc-2024-day-14/src/robot-grid.go"
)

func main() {
  fmt.Println("Solving part 1...")

  var s string = string(input.Bytes)
  var result string = strings.Trim(s, " \n")
  var lines []string = strings.Split(result, "\n")

  var robots []Robot
  for _, line := range lines {
    robot := ParseRobot(line)
    fmt.Println(robot)
    robots = append(robots, robot)
  }

  robotGrid := NewRobotGrid(robots, 11, 7)
  fmt.Println(robotGrid)
}
