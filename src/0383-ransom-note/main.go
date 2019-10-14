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

  i, o = []string{"aab", "aa"}, false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)

  i, o = []string{"bg", "abge"}, true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)

  i, o = []string{"", ""}, true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\nExpect: %t\n", canConstruct(i[0], i[1]), o)
}

// T: O(N)
// M: O(1)
// -- start --

func canConstruct(ransomNote string, magazine string) bool {
  if len(ransomNote) > len(magazine) {
    return false
  }

  count := [26]int{}

  for _, c := range magazine {
    count[c - 'a']++;
  }

  for _, c := range ransomNote {
    count[c - 'a']--;
  }

  for _, num := range count {
    if num < 0 {
      return false
    }
  }

  return true
}

// -- end --

