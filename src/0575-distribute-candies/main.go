package main
import "fmt"

func main() {
  i, o := []int{1,1,2,2,3,3}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", distributeCandies(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,1,2,3}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", distributeCandies(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(N)
// M: O(N)
// -- start --

func distributeCandies(candies []int) int {
  m := [200001]bool{}
  kinds := 0
  bound := len(candies) / 2

  for _, num := range candies {
    if m[num + 100000] {
      continue
    }

    m[num + 100000] = true
    kinds++

    if kinds > bound {
      return bound
    }
  }

  return kinds
}

// -- end --

