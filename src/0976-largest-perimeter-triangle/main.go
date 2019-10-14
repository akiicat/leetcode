package main
import "fmt"
import "sort"

func main() {
  i, o := []int{2,1,2}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", largestPerimeter(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,1}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", largestPerimeter(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{3,2,3,4}, 10
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", largestPerimeter(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{3,6,2,3}, 8
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", largestPerimeter(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n*log(n)) n for the number of A
// M: O(1)
// -- start --

func largestPerimeter(A []int) int {
  sort.Sort(sort.Reverse(sort.IntSlice(A)))  // n*log(n)

  for i := 0; i < len(A)-2; i++ {
    if A[i] < A[i+1] + A[i+2] {
      return A[i] + A[i+1] + A[i+2]
    }
  }

  return 0
}

// -- end --

