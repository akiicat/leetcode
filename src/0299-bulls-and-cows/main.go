package main
import "fmt"

func main() {
  i1, i2, o := "1807", "7810", "1A3B"
  fmt.Printf("Input:  %s %s\n", i1, i2)
  fmt.Printf("Output: %s\n", getHint(i1, i2))
  fmt.Printf("Expect: %s\n", o)

  i1, i2, o = "1123", "0111", "1A1B"
  fmt.Printf("Input:  %s %s\n", i1, i2)
  fmt.Printf("Output: %s\n", getHint(i1, i2))
  fmt.Printf("Expect: %s\n", o)

  i1, i2, o = "1", "0", "0A0B"
  fmt.Printf("Input:  %s %s\n", i1, i2)
  fmt.Printf("Output: %s\n", getHint(i1, i2))
  fmt.Printf("Expect: %s\n", o)
}

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

