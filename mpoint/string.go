package mpoint

import (
  "fmt"
)

func (p Point) String() string {
  return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}
