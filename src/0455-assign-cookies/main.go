package main
import "fmt"
import "sort"

func main() {
  i1, i2, o := []int{1,2,3}, []int{1,1}, 1
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findContentChildren(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = []int{1,2}, []int{1,2,3}, 2
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findContentChildren(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = []int{1,2}, []int{1,3}, 2
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findContentChildren(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = []int{1,2,3}, []int{}, 0
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findContentChildren(i1, i2))
  fmt.Printf("Expect: %d\n", o)

  i1, i2, o = []int{1,2,3}, []int{3}, 1
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %d\n", findContentChildren(i1, i2))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n*log(n)) for sort
// M: O(1)
// -- start --

func findContentChildren(g []int, s []int) int {
  gn, sn := len(g), len(s)
  if gn == 0 || sn == 0 {
    return 0
  }

  sort.Ints(g)
  sort.Ints(s)

  i, j, count := 0, 0, 0
  for i < gn && j < sn {
    if g[i] <= s[j] {
      count++
      i++
    }
    j++
  }

  return count
}

// -- end --

