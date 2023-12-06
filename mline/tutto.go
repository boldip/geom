package mline

import (
  "github.com/boldip/geom/mpoint"
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
