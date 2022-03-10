package main
import "fmt"
import "bytes"

// T: O(2^n)
// M: O(n)
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

