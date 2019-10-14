package main
import "fmt"
import "sort"

func main() {
  i1, i2, o := []int{1,2,3}, []int{2}, 1
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findRadius(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = []int{1,2,3,4}, []int{1,4}, 1
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findRadius(i1, i2))
  fmt.Printf("Expect: %d\n", o)
}

// T: O((m + n) * log(n)) n is the number of heaters
// M: O(1)
// -- start --

// T: O((m + n) * log(n))
// M: O(1)
func findRadius(houses []int, heaters []int) int {
  max := 0

  sort.Ints(heaters)

  for _, v := range houses {
    min := 1<<31-1
    for _, h := range heaters {
      min = Min(min, Abs(v - h))
    }
    max = Max(max, min)
  }

  return max
}

// T: O(n*m)
// M: O(1)
func findRadius(houses []int, heaters []int) int {
  max := 0

  for _, v := range houses {
    min := 1<<31-1
    for _, h := range heaters {
      min = Min(min, Abs(v - h))
    }
    max = Max(max, min)
  }

  return max
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
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

