package robot_grid

import (
	"fmt"
	"strings"

	. "valyntyler.com/aoc-2024-day-14/src/robot"
)

type RobotGrid struct {
  robots []Robot
  width int
  height int
}

func NewRobotGrid(robots []Robot, width int, height int) RobotGrid {
  return RobotGrid { robots, width, height }
}

func (r RobotGrid) grid() [][]uint {
  grid := make([][]uint, r.height)
  for i := range grid {
    grid[i] = make([]uint, r.width)
  }

  for _, robot := range r.robots {
    grid[robot.Position.Y][robot.Position.X] += 1
  }

  return grid
}

func (r RobotGrid) String() string {
  var builder strings.Builder
  grid := r.grid()

  for j := 0; j < r.height; j++ {
    for i := 0; i < r.width; i++ {
      item := grid[j][i]
      if item == 0 {
        builder.WriteString(fmt.Sprintf("%c", '.'))
      } else {
        builder.WriteString(fmt.Sprintf("%v", item))
      }
    }
    builder.WriteString("\n")
  }

  return builder.String()
}
