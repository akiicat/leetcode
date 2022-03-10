package main
import "fmt"

// T: O(n) n for length of string
// M: O(1) for 10
// -- start --

func getHint(secret string, guess string) string {
  m := [10]int{}
  a, b := 0, 0

  for i := 0; i < len(secret); i++ {
    if secret[i] == guess[i] {
      a++
    }

    if m[secret[i]-'0'] < 0 {
      b++
    }
    m[secret[i]-'0']++

    if m[guess[i]-'0'] > 0 {
      b++
    }
    m[guess[i]-'0']--
  }

  return fmt.Sprintf("%dA%dB", a, b-a)
}

// -- end --

