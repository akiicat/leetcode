package main
import "fmt"

func main() {
  s, t, o := "anagram", "nagaram", true
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isAnagram(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "rat", "car", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isAnagram(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "aa", "bb", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isAnagram(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "ac", "bb", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isAnagram(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "aacc", "bbbb", false
  fmt.Printf("Input:  %s, %s\n", s, t)
  fmt.Printf("Output: %t\n", isAnagram(s, t))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1) for 26
// -- start --

func isAnagram(s string, t string) bool {
  if len(s) != len(t) {
    return false
  }

  m := [26]int{}

  for _, v := range s {
    m[v-'a']++
  }

  for _, v := range t {
    m[v-'a']--
  }

  for _, v := range m {
    if v != 0 {
      return false
    }
  }

  return true
}

// -- end --

