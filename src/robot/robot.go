package robot

import (
	"fmt"
	"strconv"
	"strings"

	. "valyntyler.com/aoc-2024-day-14/src/vector2"
)

type Robot struct {
  Position Vector2[uint]
  Velocity Vector2[int]
}

func NewRobot(position Vector2[uint], velocity Vector2[int]) Robot {
  return Robot { position, velocity }
}

func ParseRobot(s string) Robot {
    var tokens []string = strings.Split(s, " ")

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

    return NewRobot(pos, vel)
}

func (r Robot) String() string {
  return fmt.Sprintf(
    "Robot:\n  position: %v\n  velocity: %v",
    r.Position,
    r.Velocity,
  )
}
