package main
import "fmt"

func main() {
  i, o := "PPALLP", true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", checkRecord(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "PPALLL", false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", checkRecord(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = "LALL", true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", checkRecord(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func checkRecord(s string) bool {
  a := 1
  l := 2

  for _, v := range s {
    if v == 'A' {
      if a == 0 {
        return false
      }
      a--
    }

    if v == 'L' {
      if l == 0 {
        return false
      }
      l--
    } else {
      l = 2
    }
  }

  return true
}

// -- end --
