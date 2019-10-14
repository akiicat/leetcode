package main
import "fmt"

func main() {
  i, o := 1, "A"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = 26, "Z"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = 28, "AB"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = 52, "AZ"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = 701, "ZY"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = 703, "AAA"
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %s\n", convertToTitle(i))
  fmt.Printf("Expect: %s\n", o)
}


// T: O(1)
// M: O(1)
// -- start --

func convertToTitle(n int) string {
  s := ""

  for n = n - 1; n >= 0; n = n / 26 - 1 {
    s = string(n % 26 + 65) + s
  }

  return s
}

// -- end --

