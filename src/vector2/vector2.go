package vector2

type Number interface {
  int | uint
}

type Vector2[T Number] struct {
  x T
  y T
}

func NewVector2[T Number] (x T, y T) Vector2[T] {
  return Vector2[T] { x, y }
}
