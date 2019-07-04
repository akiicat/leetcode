package main
import "fmt"
import "strconv"

func main() {
  var c []byte
  c = []byte{'a','a','b','b','c','c','c'}
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\nExpect: 6 a2b2c3\n", compress(c), c)
  c = []byte{'a'}
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\nExpect: 1 a\n", compress(c), c)
  c = []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\nExpect: 4 ab12\n", compress(c), c)
}

// T: O(N)
// M: O(1)
// -- start --

func compress(chars []byte) int {
  end, count := 0, 1

  for i, char := range chars {
    if i+1 < len(chars) && char == chars[i+1] {
      count++
    } else {
      chars[end], end = char, end + 1
      if count > 1 {
        for _, num := range strconv.Itoa(count) {
          chars[end], end = byte(num), end + 1
        }
      }
      count = 1
    }
  }

  return end
}

// -- end --

