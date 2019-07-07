package main
import "fmt"

func main() {
  fmt.Printf("Input:  123\nOutput: %d\nExpect: 321\n", reverse(123))
  fmt.Printf("Input:  -123\nOutput: %d\nExpect: -321\n", reverse(-123))
  fmt.Printf("Input:  120\nOutput: %d\nExpect: 21\n", reverse(120))
  fmt.Printf("Input:  1534236469\nOutput: %d\nExpect: 0\n", reverse(1534236469))
}

// T: O(N)
// M: O(1)
// -- start --

func reverse(x int) int {
  input := x // speed up, do not use x as computation
  if input != int(int32(input)) {
    return 0
  }

  reverse := 0
  for input != 0 {
    reverse = 10 * reverse + input % 10
    input = input / 10
  }

  if reverse != int(int32(reverse)) {
    return 0
  }

  return reverse
}

// -- end --

