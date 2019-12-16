package main
import "fmt"

func main() {
  i, o := "UD", true
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", judgeCircle(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "LL", false
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %t\n", judgeCircle(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func judgeCircle(moves string) bool {
  x, y := 0, 0
  for _, v := range moves {
    switch v {
    case 'U':
      y++
    case 'D':
      y--
    case 'R':
      x++
    case 'L':
      x--
    }
  }

  return x == 0 && y == 0
}

// -- end --

