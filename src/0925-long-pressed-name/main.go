package main
import "fmt"

func main() {
  name, typed, o := "alex", "aaleex", true
  fmt.Printf("Input:  %s %s\n", name, typed)
  fmt.Printf("Output: %t\n", isLongPressedName(name, typed))
  fmt.Printf("Expect: %t\n", o)

  name, typed, o = "saeed", "ssaaedd", false
  fmt.Printf("Input:  %s %s\n", name, typed)
  fmt.Printf("Output: %t\n", isLongPressedName(name, typed))
  fmt.Printf("Expect: %t\n", o)

  name, typed, o = "leelee", "lleeelee", true
  fmt.Printf("Input:  %s %s\n", name, typed)
  fmt.Printf("Output: %t\n", isLongPressedName(name, typed))
  fmt.Printf("Expect: %t\n", o)

  name, typed, o = "vtkgn", "vttkgnn", true
  fmt.Printf("Input:  %s %s\n", name, typed)
  fmt.Printf("Output: %t\n", isLongPressedName(name, typed))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n) n is the length of typed
// M: O(1)
// -- start --

func isLongPressedName(name string, typed string) bool {
  i := 0
  for j, _ := range typed {
    if name[i] != typed[j] {
      continue
    }

    i++
    if i == len(name) {
      return true
    }
  }

  return i == len(name)
}

// -- end --

