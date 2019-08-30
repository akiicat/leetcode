package main
import "fmt"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := 38, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", addDigits(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(1)
// M: O(1)
// -- start --

func addDigits(num int) int {
  // repeat every 9 numbers except 0
  // 1  2  3  4  5  6  7  8  9
  // 10 11 12 13 14 15 16 17 18
  // 19 20 21 22 23 24 25 26 27

  if num == 0 {
    return 0
  }

  if num % 9 == 0 {
    return 9
  }

  return num % 9
}

func addDigitsRecursive(num int) int {
  if num < 10 {
    return num
  }

  sum := 0
  for num > 0 {
    sum += num % 10
    num /= 10
  }

  return addDigits(sum)
}

// -- end --

