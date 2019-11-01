package main
import "fmt"
import "strconv"

func main() {
  i, o := 100, "202"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", convertToBase7(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = 99, "201"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", convertToBase7(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = -7, "-10"
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", convertToBase7(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

// Implement `strconv.FormatInt(int64(num), 7)` function
func convertToBase7(num int) string {
  if num == 0 {
    return "0"
  }

  p := ""
  if num < 0 {
    num = -num
    p = "-"
  }

  s := ""
  for num != 0 {
    s = strconv.Itoa(num % 7) + s
    num /= 7
  }

  return p + s
}

// -- end --

