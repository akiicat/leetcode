package main
import "fmt"
import "strings"

func main() {
  a, b, o := "this apple is sweet", "this apple is sour", []string{"sweet", "sour"}
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %v\n", uncommonFromSentences(a, b))
  fmt.Printf("Expect: %v\n", o)

  a, b, o = "apple apple", "banana", []string{"banana"}
  fmt.Printf("Input:  %s %s\n", a, b)
  fmt.Printf("Output: %v\n", uncommonFromSentences(a, b))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n + m) n, m are the lengths of A, B respectively
// M: O(n + m)
// -- start --

func uncommonFromSentences(A string, B string) []string {
  m := make(map[string]int)

  for _, v := range strings.Split(A, " ") {
    m[v]++
  }
  for _, v := range strings.Split(B, " ") {
    m[v]++
  }

  rtn := []string{}
  for k, v := range m {
    if v == 1 {
      rtn = append(rtn, k)
    }
  }

  return rtn
}

// -- end --

