package main
import "fmt"
import "strconv"

func main() {
  c, o1, o2 := []byte{'a','a','b','b','c','c','c'}, 6, "a2b2c3"
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\n", compress(c), c[:o1])
  fmt.Printf("Expect: %d %s\n", o1, o2)

  c, o1, o2 = []byte{'a'}, 1, "a"
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\n", compress(c), c[:o1])
  fmt.Printf("Expect: %d %s\n", o1, o2)

  c, o1, o2 = []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}, 4, "ab12"
  fmt.Printf("Input:  %s\n", c)
  fmt.Printf("Output: %d %s\n", compress(c), c[:o1])
  fmt.Printf("Expect: %d %s\n", o1, o2)
}

// T: O(n)
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

