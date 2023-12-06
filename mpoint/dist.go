package mpoint

import (
  "math"
)

func (p1 Point) Dist(p2 Point) float64 {
  dx := p1.x - p2.x
  dy := p1.y - p2.y
  return math.Sqrt(dx * dx + dy * dy)
}
