package main
import "fmt"

func main() {
  i, o := []string{"bella","label","roller"}, []string{"e","l","l"}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", commonChars(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []string{"cool","lock","cook"}, []string{"c","o"}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", commonChars(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(N) N for all of chars
// M: O(1)
// -- start --

func commonChars(A []string) []string {
  m := make([]int, 26)

  for i := 0; i < 26; i++ {
    m[i] = (1 << 31) - 1
  }

  for _, str := range A {
    tmp := make([]int, 26)
    for _, c := range []byte(str) {
      tmp[c - 'a']++
    }

    for k, num := range tmp {
      m[k] = min(num, m[k])
    }
  }

  rtn := []string{}

  for i := 0; i < 26; i++ {
    if m[i] == (1 << 31) - 1 {
      continue
    }

    for j := 0; j < m[i]; j++ {
      rtn = append(rtn, string(byte(i+'a')))
    }
  }

  return rtn
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

