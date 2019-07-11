package main
import "fmt"

func main() {
  i, o := []string{"a", "b"}, false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)

  i, o = []string{"aa", "ab"}, false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)

  i, o = []string{"aa", "aab"}, true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)
}

// T: O(N)
// M: O(1)
// -- start --

func canConstruct(ransomNote string, magazine string) bool {
  len_l := len(ransomNote)
  return false
}

// -- end --

