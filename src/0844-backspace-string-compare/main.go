package main
import "fmt"

func main() {
  s, t, o := "ab#c", "ad#c", true
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", backspaceCompare(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "ab##", "c#d#", true
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", backspaceCompare(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "a##c", "#a#c", true
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", backspaceCompare(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "a#c", "b", false
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", backspaceCompare(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = "y#fo##f", "y#f#o##f", true
  fmt.Printf("Input:  %s %s\n", s, t)
  fmt.Printf("Output: %t\n", backspaceCompare(s, t))
  fmt.Printf("Expect: %t\n", o)

}

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func backspaceCompare(S string, T string) bool {
  n, m := len(S), len(T)

  for {
    n, m = helper(S, n-1), helper(T, m-1)
    if n < 0 || m < 0 || S[n] != T[m] {
      return n == -1 && m == -1
    }
  }

  return false
}

// T: O(n + m)
// M: O(1)
func backspaceCompareTwoPointer(S string, T string) bool {
  n, m := len(S), len(T)

  for n >= 0 && m >= 0 {
    n, m = helper(S, n-1), helper(T, m-1)
    if n >= 0 && m >= 0 && S[n] != T[m] {
      return false
    }
  }

  return n == -1 && m == -1
}

func helper(s string, i int) int {
  skip := 0

  for i >=0 {
    if s[i] == '#' {
      skip++
    } else if skip > 0 {
      skip--
    } else {
      return i
    }

    i--
  }

  return i // -1
}

// T: O(n + m)
// M: O(max(n, m))
func backspaceCompareStack(S string, T string) bool {
  return typing(S) == typing(T)
}

func typing(S string) string {
  s := []rune{} // stack

  for _, v := range S {
    if v != '#' {
      s = append(s, v)
    } else if len(s) != 0 {
      s = s[:len(s)-1]
    }
  }

  return string(s)
}

// -- end --

