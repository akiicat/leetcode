package main
import "fmt"
import "math"
import "strings"

func main() {
  i, o := "abab", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "aba", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abcabcabcabc", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "ababba", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "aabaaba", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "bb", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "ababababab", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "abaababaab", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "ab", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", repeatedSubstringPattern(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(log(N))
// M: O(1)
// -- start --

func repeatedSubstringPattern(s string) bool {
	ss := (s + s)[1:len(s)*2-1] // shift one
	return strings.Contains(ss, s)
}

func repeatedSubstringPatternRepeat(s string) bool {
	n := len(s)

  upperBound := int(math.Sqrt((float64(n)))) + 1

	for i := 1; i < upperBound; i ++ {
		if n % i != 0 {
      continue
		}

    if repeat(s, i, n / i) {
      return true
    }
	}

	return false
}

func repeat(s string, p, q int) bool {
  // return repeat1(s, p, q) || repeat1(s, q, p)
  return repeat2(s, p, q) || repeat2(s, q, p)
}

func repeat1(s string, p, q int) bool {
  return q > 1 && strings.Repeat(s[:p], q) == s
}

func repeat2(s string, p, q int) bool {
  for i := 1; i < q; i++ {
    if s[:p] != s[i*p:(i+1)*p] {
      return false
    }
  }

  return q > 1
}

// -- end --

