package vector2

import "fmt"

type Number interface {
  int | uint
}

type Vector2[T Number] struct {
  X T
  Y T
}

func NewVector2[T Number] (x T, y T) Vector2[T] {
  return Vector2[T] { x, y }
}

func (v Vector2[T]) String() string {
  return fmt.Sprintf("Vector2(%d, %d)", v.X, v.Y)
}
