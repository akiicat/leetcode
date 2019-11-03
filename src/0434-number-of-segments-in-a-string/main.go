package main
import "fmt"

func main() {
  i, o := "Hello, my name is John", 5
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", countSegments(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "Hello", 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", countSegments(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "", 0
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", countSegments(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "           ", 0
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", countSegments(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n) n is the length of string
// M: O(1)
// -- start --

// implement len(strings.Split(s, " "))
func countSegments(s string) int {
  c := 0

  for i, v := range s {
    if (i == 0 || s[i-1] == ' ') && v != ' ' {
      c++
    }
  }

  return c
}

func countSegmentsCountWords(s string) int {
  w, c := 0, 0

  for _, v := range s {
    if v == ' ' {
      if w != 0 {
        c++
      }
      w = 0
    } else {
      w++
    }
  }

  if w != 0 {
    c++
  }

  return c
}

// -- end --

