package main
import "fmt"

func main() {
  i, o := "leetcode", 0
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", firstUniqChar(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "loveleetcode", 2
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", firstUniqChar(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "lel", 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", firstUniqChar(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "cc", -1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", firstUniqChar(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func firstUniqChar(s string) int {
  m := [26]int{}

  for i := 0; i < len(s); i++ {
    m[s[i]-'a']++
  }

  for i := 0; i < len(s); i++ {
    if m[s[i]-'a'] == 1 {
      return i
    }
  }

  return -1
}

func firstUniqCharDiscreteTransform(s string) int {
  m := [26]int{}

  for i := 0; i < 26; i++ {
    m[i] = -1
  }

  n := len(s)
  for i := 0; i < n; i++ {
    if m[s[i]-'a'] == -1 {
      m[s[i]-'a'] = i
    } else {
      m[s[i]-'a'] = n
    }
  }

  first := n
  for i := 0; i < 26; i++ {
    if m[i] >= 0 && m[i] < first {
      first = m[i]
    }
  }

  if first == n {
    return -1
  }

  return first
}

// -- end --
