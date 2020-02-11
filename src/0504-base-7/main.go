package main
import "strconv"

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

