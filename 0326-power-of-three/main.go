package main
import "fmt"
import "math"

func main() {
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 0, isPowerOfThree(0), false)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 1, isPowerOfThree(1), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 3, isPowerOfThree(3), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 9, isPowerOfThree(9), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 27, isPowerOfThree(27), true)
  fmt.Printf("Input:  %d\nOutput: %t\nExpect: %t\n", 45, isPowerOfThree(45), false)
}

// T: O(1)
// M: O(1)
// -- start --

func isPowerOfThree(n int) bool {
  // int(math.Pow(3, 20)) == 3486784401
  // int(math.Pow(2, 31)) == 2147483648
  return n > 0 && 3486784401 % n == 0
}

func isPowerOfThreeLog(n int) bool {
  res := math.Log(float64(n)) / math.Log(3)
  return (res - float64(int(res))) < 1E-14
}

func isPowerOfThreeForLoop(n int) bool {
  if n <= 0 {
    return false
  }

  for n != 1 {
    if n % 3 != 0 {
      return false
    }
    n /= 3
  }

  return true
}

// -- end --

