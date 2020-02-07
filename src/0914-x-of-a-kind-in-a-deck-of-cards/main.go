package main
import "fmt"

func main() {
  i, o := []int{1,2,3,4,4,3,2,1}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1,1,2,2,2,3,3}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{0,0,0,1,1,1,2,2,2}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1,2,2,2,2}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1,1,1,2,2,2,2,2,2}, true
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = []int{1,1,1,1,1,1,2,2,2,2,2,2,2,2,2,3,3,3,3,3,3,3,3}, false
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %t\n", hasGroupsSizeX(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n * log(n) ** 2)
// M: O(n)
// -- start --

func hasGroupsSizeX(deck []int) bool {
  m := make(map[int]int)
  for _, v := range deck {
    m[v]++
  }

  t := -1
  for _, v := range m {
    if t == -1 {
      t = v
    } else {
      t = gcd(t, v)
    }
  }

  return t >= 2
}

func gcd(a, b int) int {
  for b != 0 {
    a, b = b, a % b
  }
  return a
}

// -- end --
