package robot

import (
	"fmt"

	. "valyntyler.com/aoc-2024-day-14/src/vector2"
)

type Robot struct {
  position Vector2[uint]
  velocity Vector2[int]
}

func NewRobot(position Vector2[uint], velocity Vector2[int]) Robot {
  return Robot { position, velocity }
}

func (r Robot) String() string {
  return fmt.Sprintf(
    "Robot:\n  position: %v\n  velocity: %v",
    r.position,
    r.velocity,
  )
}
