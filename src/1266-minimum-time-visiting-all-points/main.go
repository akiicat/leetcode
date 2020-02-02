package main
import "fmt"

func main() {
  i, o := [][]int{[]int{1,1}, []int{3,4}, []int{-1,0}}, 7
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minTimeToVisitAllPoints(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = [][]int{[]int{3,2}, []int{-2,2}}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", minTimeToVisitAllPoints(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n) n is the number of points
// M: O(1)
// -- start --

func minTimeToVisitAllPoints(points [][]int) int {
  sum := 0

  for i := 0; i < len(points)-1; i++ {
    dx := Abs(points[i][0] - points[i+1][0])
    dy := Abs(points[i][1] - points[i+1][1])

    sum += Max(dx, dy)
  }

  return sum
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// -- end --

