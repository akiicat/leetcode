package main
import "fmt"

func main() {
  s, k, o := "5F3Z-2e-9-w", 4, "5F3Z-2E9W"
  fmt.Printf("Input:  %s %d\n", s, k)
  fmt.Printf("Output: %s\n", licenseKeyFormatting(s, k))
  fmt.Printf("Expect: %s\n", o)

  s, k, o = "2-5g-3-J", 2, "2-5G-3J"
  fmt.Printf("Input:  %s %d\n", s, k)
  fmt.Printf("Output: %s\n", licenseKeyFormatting(s, k))
  fmt.Printf("Expect: %s\n", o)

  s, k, o = "--a-a-a-a--", 2, "AA-AA"
  fmt.Printf("Input:  %s %d\n", s, k)
  fmt.Printf("Output: %s\n", licenseKeyFormatting(s, k))
  fmt.Printf("Expect: %s\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func licenseKeyFormattingStack(S string, K int) string {
  c := 0
  s := []byte{} // stack

  for i := len(S)-1; i >= 0; i-- {
    if S[i] == '-' {
      continue
    }

    if c == K {
      s, c = append(s, '-'), 0
    }

    if S[i] >= 97 {
      s, c = append(s, S[i]-32), c + 1
    } else {
      s, c = append(s, S[i]), c + 1
    }
  }

  // reverse stack
  l, r := 0, len(s)-1
  for l < r {
    s[l], s[r] = s[r], s[l]
    l++
    r--
  }

  return string(s)
}

// -- end --

