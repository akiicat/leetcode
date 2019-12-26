package main
import "fmt"

func main() {
  i, o := "III", 3
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", romanToInt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "IV", 4
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", romanToInt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "IX", 9
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", romanToInt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "LVIII", 58
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", romanToInt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "MCMXCIV", 1994
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", romanToInt(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n) n is the sum of all characters in all strings.
// M: O(1)
// -- start --

func romanToInt(s string) int {
  m := map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  }

  sum, pre := 0, 0

  for i := len(s) - 1; i >= 0; i-- {
    cur := m[s[i]]
    if cur >= pre {
      sum += cur
    } else {
      sum -= cur
    }
    pre = cur
  }

  return sum
}

// -- end --

