package main
import "fmt"
import "math"
import "strconv"

func main() {
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 0, isPowerOfFour(0), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 1, isPowerOfFour(1), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 2, isPowerOfFour(2), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 5, isPowerOfFour(5), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 16, isPowerOfFour(16), true)
}

// T: O(1)
// M: O(1)
// -- start --

func isPowerOfFour(num int) bool {
  if num <= 0 {
    return false
  }

  s := strconv.FormatInt(int64(num), 2)
  length := len(s)
  shiftValue, _ := strconv.Atoi(s[1:])

  return shiftValue == 0 && length % 2 == 1
}

func isPowerOfFourLog(num int) bool {
  res := math.Log(float64(num)) / math.Log(4)
  return (res - float64(int(res))) < 1E-14
}

func isPowerOfFourForLoop(n int) bool {
  if n <= 0 {
    return false
  }

  for n != 1 {
    if n % 4 != 0 {
      return false
    }
    n >>= 2
  }

  return true
}

// -- end --

