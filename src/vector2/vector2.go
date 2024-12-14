package vector2

type Number interface {
  int | uint
}

type Vector2[T Number] struct {
  x T
  y T
}
