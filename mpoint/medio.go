package mpoint

func (p1 Point) Median(p2 Point) Point {
  var res Point
  res = NewPoint((p1.x + p2.x) / 2, (p1.y+p2.y) / 2)
  return res
}
