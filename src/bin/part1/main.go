package main

import (
	"fmt"
	"strings"
	"strconv"

	"valyntyler.com/aoc-2024-day-14/input"
	. "valyntyler.com/aoc-2024-day-14/src/vector2"
	. "valyntyler.com/aoc-2024-day-14/src/robot"
)

func main() {
  fmt.Println("Solving part 1...")

  // var robots []Robot

  var s string = string(input.Bytes)
  var result string = strings.Trim(s, " \n")
  var lines []string = strings.Split(result, "\n")

  for _, line := range lines {
    var tokens []string = strings.Split(line, " ")

    var pos_string = strings.Trim(tokens[0], "p=")
    var vel_string = strings.Trim(tokens[1], "v=")

    var pos_values = strings.Split(pos_string, ",")
    pos_x, _ := strconv.ParseUint(pos_values[0], 10, 32)
    pos_y, _ := strconv.ParseUint(pos_values[1], 10, 32)
    pos := NewVector2(uint(pos_x), uint(pos_y))

    var vel_values = strings.Split(vel_string, ",")
    vel_x, _ := strconv.ParseInt(vel_values[0], 10, 32)
    vel_y, _ := strconv.ParseInt(vel_values[1], 10, 32)
    vel := NewVector2(int(vel_x), int(vel_y))

    robot := NewRobot(pos, vel)
    fmt.Println(robot)
  }
}
