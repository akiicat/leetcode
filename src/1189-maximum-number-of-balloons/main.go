package main
import "fmt"

func main() {
  i, o := "nlaebolko", 1
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", maxNumberOfBalloons(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "loonbalxballpoon", 2
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", maxNumberOfBalloons(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = "leetcode", 0
  fmt.Printf("Input:  %s\n", i)
  fmt.Printf("Output: %d\n", maxNumberOfBalloons(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1) 26 chars
// -- start --

func maxNumberOfBalloons(text string) int {
  b := map[rune]int{
    'b': 1,
    'o': 2,
    'l': 2,
    'a': 1,
    'n': 1,
  }

  m := make(map[rune]int)

  for _, v := range text {
    m[v]++
  }

  min := 1<<31
  for k, v := range b {
    min = Min(min, m[k] / v)
  }

  return min
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

