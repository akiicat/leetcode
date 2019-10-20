package main
import "fmt"

func main() {
  i, o := 5, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", findComplement(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 6, 1
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", findComplement(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 1, 0
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", findComplement(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(1) max 32 times
// M: O(1)
// -- start --

func findComplement(num int) int {
  mask := ^0     // 0xffffffff = -1

  for num & mask > 0 {
    mask <<= 1
  }

  return num ^ (^mask)
}

func findComplementShiftLeft(num int) int {
  if num == 0 {
    return 1
  }

  t, c := num, uint(0)

  for t != 0 {
    t = t >> 1
    c++
  }

  return num ^ (1 << c - 1)
}

// -- end --

