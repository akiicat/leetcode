package main

// T: O(n)
// M: O(1)
// -- start --

func checkStraightLine(coordinates [][]int) bool {
  x1, y1 := coordinates[0][0], coordinates[0][1]
  x2, y2 := coordinates[1][0], coordinates[1][1]

  dx1, dy1 := x2 - x1, y2 - y1

  for _, v := range coordinates[2:] {
    dx2, dy2 := v[0] - x1, v[1] - y1

    if dx1 * dy2 != dx2 * dy1 {
      return false
    }
  }

  return true
}

func checkStraightLineDouble(coordinates [][]int) bool {
  x, y := coordinates[0][0], coordinates[0][1]
  m := float64(coordinates[1][1] - y) / float64(coordinates[1][0] - x)

  for _, v := range coordinates[2:] {
    if m != float64(v[1] - y) / float64(v[0] - x) {
      return false
    }
  }

  return true
}

// -- end --

