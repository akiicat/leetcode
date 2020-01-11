package main
import "fmt"
import "strings"

func main() {
  p, b, o := "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"
  fmt.Printf("Input:  %s, %v\n", p, b)
  fmt.Printf("Output: %s\n", mostCommonWord(p, b))
  fmt.Printf("Expect: %s\n", o)

  p, b, o = "Bob", []string{}, "bob"
  fmt.Printf("Input:  %s, %v\n", p, b)
  fmt.Printf("Output: %s\n", mostCommonWord(p, b))
  fmt.Printf("Expect: %s\n", o)
}

// T: O(n + m)
// M: O(n + m)
// -- start --

func mostCommonWord(paragraph string, banned []string) string {
  ban := make(map[string]bool)
  for _, v := range banned {
    ban[v] = true
  }

  m := make(map[string]int)
  l := -1
  paragraph = strings.ToLower(paragraph) + " "
  for i, v := range paragraph {
    if 'a' <= v && v <= 'z' {
      continue
    }

    w := paragraph[l+1:i]
    if len(w) > 0 && !ban[w] {
      m[w]++
    }

    l = i
  }

  rtn, max := "", 0
  for k, v := range m {
    if v > max {
      rtn, max = k, v
    }
  }

  return rtn
}

// -- end --

