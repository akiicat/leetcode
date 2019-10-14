package main
import "fmt"
import "math/bits"

func main() {
  fmt.Printf("Input:  1, 4\nOutput: %d\nExpect: 2\n", hammingDistance(1, 4))
  fmt.Printf("Input:  93, 73\nOutput: %d\nExpect: 2\n", hammingDistance(93, 73))
}

// T: O(1)
// M: O(1)
// -- start --

func hammingDistance(x int, y int) int {
  return bits.OnesCount(uint(x ^ y))
}

// -- end --

