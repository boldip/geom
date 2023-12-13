package point

/*
go test
go test -v

go test -coverprofile cover.out
go tool cover -html=cover.out
*/


import (
  "testing"
  "math"
)

func almostEqual(a, b float64) bool {
  return math.Abs(a-b) < 1E-6
}

func TestMedian(t *testing.T) {
  var p1, p2, pm Point
  p1 = NewPoint(0, 0)
  p2 = NewPoint(2, 4)
  pm = Median(p1, p2)

  if !( almostEqual(pm.X, 1) && almostEqual(pm.Y, 2) ) {
    t.Error("Expected (1,2), got ", pm)
  }
}

func TestDist(t *testing.T) {
  var p1, p2 Point
  p1 = NewPoint(0, 0)
  p2 = NewPoint(2, 4)
  actual := Dist(p1, p2)
  expected := math.Sqrt(2*2 + 4*4)
  if !( almostEqual(actual, expected) ) {
    t.Error("Expected ", expected, ", got ", actual)
  }

}
