package main

// T: O(n*3)
// M: O(1)
// -- start --

func largestTriangleArea(points [][]int) float64 {
  if len(points) < 3 {
    return 0
  }

  max := 0.0

  for i := 0; i < len(points)-2; i++ {
    for j := 0; j < len(points)-1; j++ {
      for k := 0; k < len(points); k++ {
        area := Area(points[i], points[j], points[k])

        if area > max {
          max = area
        }
      }
    }
  }

  return max / 2
}

func Area(p1, p2, p3 []int) float64 {
  a0 := p2[0] - p1[0]
  a1 := p2[1] - p1[1]
  b0 := p3[0] - p1[0]
  b1 := p3[1] - p1[1]

  return float64(Abs(a0 * b1 - a1 * b0))
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// -- end --

