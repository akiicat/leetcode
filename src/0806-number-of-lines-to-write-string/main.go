package main
import "fmt"

func main() {
  widths, S, o := []int{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}, "abcdefghijklmnopqrstuvwxyz", []int{3,60}
  fmt.Printf("Input:  %v, %s\n", widths, S)
  fmt.Printf("Output: %v\n", numberOfLines(widths, S))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n is the length of S
// M: O(1)
// -- start --

func numberOfLines(widths []int, S string) []int {
  line, index := 1, 0

  for _, v := range S {
    width := widths[v-'a']

    if index + width > 100 {
      line, index = line+1, 0 // reset
    }

    index += width
  }

  return []int{line, index}
}

// -- end --

