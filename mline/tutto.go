package mline

import (
  "github.com/boldip/geom/mpoint"
  "fmt"
  "math"
)

/* A (straight) line specified by the equation y=mx+q */
type Line struct {
  m, q float64
}

func NewLine(m, q float64) (res Line) {
  res.m,  res.q = m, q
  return
}

func (r1 Line) IsParallel(r2 Line) bool {
  return r1.m == r2.m
}

func (r Line) Belongs(p mpoint.Point) bool {
  return p.Y() == r.m * p.X() + r.q
}

func (r Line) String() string {
  return fmt.Sprintf("y=%.2f x + %.2f", r.m, r.q)
}

func (r1 Line) Intersection(r2 Line) (mpoint.Point, bool) {
  if r1.m == r2.m  {
    return mpoint.NewPoint(0,0), false
  }
  x := (r2.q - r1.q) / (r1.m - r2.m)
  y := r1.m * x + r1.q
  return mpoint.NewPoint(x,y), true
}

func (r Line) Dist(p mpoint.Point) float64 {
  return math.Abs(r.q + r.m * p.X() - p.Y()) / math.Sqrt(1 + r.m * r.m)
}
