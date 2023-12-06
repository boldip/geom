package mpoint

func NewPoint(x, y float64) Point {
  var p Point
  p.x, p.y = x, y
  return p
}
