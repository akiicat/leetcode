package main
import "fmt"
import "strings"

func main() {
  i, first, second, o := "alice is a good girl she is a good student", "a", "good", []string{"girl", "student"}
  fmt.Printf("Input:  %s, %s, %s\n", i, first, second)
  fmt.Printf("Output: %v\n", findOcurrences(i, first, second))
  fmt.Printf("Expect: %v\n", o)

  i, first, second, o = "we will we will rock you", "we", "will", []string{"we", "rock"}
  fmt.Printf("Input:  %s, %s, %s\n", i, first, second)
  fmt.Printf("Output: %v\n", findOcurrences(i, first, second))
  fmt.Printf("Expect: %v\n", o)

  i, first, second, o = "a a a a a a alice is a good girl she is a good student", "a", "a", []string{"a","a","a","a","alice"}
  fmt.Printf("Input:  %s, %s, %s\n", i, first, second)
  fmt.Printf("Output: %v\n", findOcurrences(i, first, second))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n is the number of words
// M: O(1)
// -- start --

func findOcurrences(text string, first string, second string) []string {
  t := strings.Split(text, " ")

  if len(t) < 3 {
    return []string{}
  }

  rtn := []string{}

  for i := 0; i < len(t) - 2; i++ {
    if t[i] == first && t[i+1] == second {
      rtn = append(rtn, t[i+2])
    }
  }

  return rtn
}

// -- end --

