package main
import "fmt"

func main() {
  i, o := "A", 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "Z", 26
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "AB", 28
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "AZ", 52
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "ZY", 701
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "AAA", 703
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", titleToNumber(i))
  fmt.Printf("Expect: %d\n", o)
}


// T: O(1)
// M: O(1)
// -- start --

func titleToNumber(s string) int {
  n := 0
  for _, v := range s {
    n = 26 * n + int(v) - 64
  }
  return n
}

func convertToTitle(n int) string {
  s := ""

  for n = n - 1; n >= 0; n = n / 26 - 1 {
    s = string(n % 26 + 65) + s
  }

  return s
}

// -- end --

