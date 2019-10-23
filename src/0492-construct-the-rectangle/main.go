package main
import "fmt"
import "math"

func main() {
  i, o := 4, []int{2,2}
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %v\n", constructRectangle(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n ^ 1/2)
// M: O(1)
// -- start --

func constructRectangle(area int) []int {
  w := int(math.Sqrt(float64(area)))

  for w > 1 {
    if area % w == 0 {
      return []int{area/w, w}
    }
    w--
  }

  return []int{area,1}
}

// -- end --

