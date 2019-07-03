package main
import "fmt"
import "bytes"

func main() {
  fmt.Printf("Input:  1\nOutput: %s\nExpect: 1\n", countAndSay(1))
  fmt.Printf("Input:  2\nOutput: %s\nExpect: 11\n", countAndSay(2))
  fmt.Printf("Input:  3\nOutput: %s\nExpect: 21\n", countAndSay(3))
  fmt.Printf("Input:  4\nOutput: %s\nExpect: 1211\n", countAndSay(4))
  fmt.Printf("Input:  5\nOutput: %s\nExpect: 111221\n", countAndSay(5))
}

// T: O(2^N)
// M: O(N)
// -- start --

var initialize = false
var m = make(map[int]string)

func countAndSay(n int) string {
  if !initialize {
    m[1] = "1"
    for i := 2; i <= 30; i++ {
      m[i] = next(m[i-1])
    }
    initialize = true
  }

  return m[n]
}

func next(s string) string {
  c, m := 1, s[0]
  rtn := new(bytes.Buffer)

  for i := 1; i < len(s); i++ {
    if s[i] == m {
      c++
      continue
    }
    fmt.Fprintf(rtn, "%d%c", c, m)
    c, m = 1, s[i]
  }

  fmt.Fprintf(rtn, "%d%c", c, m)

  return rtn.String()
}

func countAndSayRecursive(n int) string {
  if n == 1 {
    return "1"
  }

  s := countAndSayRecursive(n-1)

  c, m := 1, s[0]
  rtn := new(bytes.Buffer)
  for i := 1; i < len(s); i++ {
    if s[i] == m {
      c++
      continue
    }
    fmt.Fprintf(rtn, "%d%c", c, m)
    c, m = 1, s[i]
  }
  fmt.Fprintf(rtn, "%d%c", c, m)

  return rtn.String()
}

// -- end --

