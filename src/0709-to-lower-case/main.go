package main
import "fmt"

func main() {
  i, o := "Hello", "hello"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", toLowerCase(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = "here", "here"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", toLowerCase(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = "LOVELY", "lovely"
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %s\n", toLowerCase(i))
  fmt.Printf("Expect: %s\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func toLowerCase(str string) string {
  rtn := []rune{}
  for _, v := range str {
    if 'A' <= v && v <= 'Z' {
      v = v - 'A' + 'a'
    }
    rtn = append(rtn, v)
  }
  return string(rtn)
}

// -- end --

