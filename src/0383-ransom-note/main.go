package main
import "fmt"

func main() {
  r, m, o := "a", "b", false
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)

  r, m, o = "aa", "ab", false
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)

  r, m, o = "aa", "aab", true
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)

  r, m, o = "aab", "aa", false
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)

  r, m, o = "bg", "abge", true
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)

  r, m, o = "", "", true
  fmt.Printf("Input:  %s, %s\n", r, m)
  fmt.Printf("Output: %t\n", canConstruct(r, m))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
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

