package main
import "fmt"

func main() {
  l, r, o := 1, 22, []int{1,2,3,4,5,6,7,8,9,11,12,15,22}
  fmt.Printf("Input:  %d %d\n", l, r)
  fmt.Printf("Output: %v\n", selfDividingNumbers(l, r))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n is the number of integers in the range
// M: O(n) answer
// -- start --

func selfDividingNumbers(left int, right int) []int {
  m := []int{}

  for i := left; i <= right; i++ {
    if Div(i) {
      m = append(m, i)
    }
  }

  return m
}

func Div(i int) bool {
  if i < 10 {
    return true
  }

  n := i

  for n != 0 {
    mod := n % 10
    if mod == 0 || i % mod != 0 {
      return false
    }
    n /= 10
  }

  return true
}

// -- end --

