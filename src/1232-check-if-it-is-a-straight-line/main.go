package main
import "fmt"

func main() {
  i, o := [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,5},[]int{5,6},[]int{6,7}}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", checkStraightLine(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = [][]int{[]int{1,1},[]int{2,2},[]int{3,4},[]int{4,5},[]int{5,6},[]int{7,7}}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", checkStraightLine(i))
  fmt.Printf("Expect: %t\n", o)
}

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

