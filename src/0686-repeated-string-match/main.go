package main
import "fmt"
import "strings"

func main() {
  i, o := []string{"abcd", "cdabcdab"}, 3
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"abcd", "cdab"}, 2
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"a", "aa"}, 2
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"aa", "aaa"}, 2
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"abc", "cabcabca"}, 4
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"ab", "bababababa"}, 6
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"abcabcabcabc", "abac"}, -1
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)

  i, o = []string{"aa", "a"}, 1
  fmt.Printf("Input:  %s, %s\n", i[0], i[1])
  fmt.Printf("Output: %d\n", repeatedStringMatch(i[0], i[1]))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n + m)
// M: O(1)
// -- start --

// T: O(n + m)
// M: O(1)
func repeatedStringMatch(A string, B string) int {
  bound, cur, str := len(B) / len(A) + 2, 1, A

  for cur <= bound {
    if strings.Index(A, B) != -1 {
      return cur
    }

    cur, A = cur + 1, A + str
  }

  return -1
}

// -- end --

