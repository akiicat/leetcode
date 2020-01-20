package main
import "fmt"

func main() {
  r1, r2, o := []int{0,0,2,2}, []int{1,1,3,3}, true
  fmt.Printf("Input:  %v %v\n", r1, r2)
  fmt.Printf("Output: %t\n", isRectangleOverlap(r1, r2))
  fmt.Printf("Expect: %t\n", o)

  r1, r2, o = []int{0,0,1,1}, []int{1,0,2,1}, false
  fmt.Printf("Input:  %v %v\n", r1, r2)
  fmt.Printf("Output: %t\n", isRectangleOverlap(r1, r2))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(1)
// M: O(1)
// -- start --

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
  return !(rec1[2] <= rec2[0] ||
           rec1[3] <= rec2[1] ||
           rec1[0] >= rec2[2] ||
           rec1[1] >= rec2[3])
}

// -- end --

