package main
import "fmt"

func main() {
  a, b, o := "ab", "ba", true
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", buddyStrings(a, b))
  fmt.Printf("Expect: %t\n", o)

  a, b, o = "ab", "ab", false
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", buddyStrings(a, b))
  fmt.Printf("Expect: %t\n", o)

  a, b, o = "aa", "aa", true
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", buddyStrings(a, b))
  fmt.Printf("Expect: %t\n", o)

  a, b, o = "abcd", "badc", false
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", buddyStrings(a, b))
  fmt.Printf("Expect: %t\n", o)

  a, b, o = "aaaaaaabc", "aaaaaaacb", true
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %t\n", buddyStrings(a, b))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1) for 26
// -- start --

// T: O(n)
// M: O(1)
func buddyStrings(A string, B string) bool {
  if len(A) != len(B) {
    return false
  }

  m := [26]int{}
  d := []int{}

  for i := 0; i < len(A); i++ {
    if A[i] != B[i] {
      d = append(d, i)
    }

    if len(d) > 2 {
      return false
    }

    if A[i] == B[i] {
      m[A[i]-'a']++
    }
  }

  if A != B {
    return A[d[0]] == B[d[1]] && A[d[1]] == B[d[0]]
  }

  for _, v := range m {
    if v > 1 {
      return true
    }
  }

  return false
}

// T: O(n)
// M: O(1)
func buddyStringsIntuition(A string, B string) bool {
  if len(A) != len(B) {
    return false
  }

  m := [26]int{}

  if A == B {
    for i := 0; i < len(A); i++ {
      m[A[i]-'a']++
      if m[A[i]-'a'] >= 2 {
        return true
      }
    }

    return false
  }

  c := 0

  for i := 0; i < len(A); i++ {
    if A[i] != B[i] {
      m[c] = i
      c++
    }

    if c > 2 {
      return false
    }
  }

  return A[m[0]] == B[m[1]] && A[m[1]] == B[m[0]]
}

// -- end --

