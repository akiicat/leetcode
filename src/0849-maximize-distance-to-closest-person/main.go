package main
import "fmt"

func main() {
  i, o := []int{1,0,0,0,1,0,1}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxDistToClosest(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,0,0,0,0,1,0,1}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxDistToClosest(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,0,0,0}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxDistToClosest(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{0,0,1,0}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", maxDistToClosest(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func maxDistToClosest(seats []int) int {

  side := true
  count := 0
  max := 0

  for _, v := range seats {
    if v == 0 {
      count++
      continue
    }

    if side == true {
      side = false
      max = count
      count = 0
      continue
    }

    max = Max(max, (count + 1) / 2)
    count = 0
  }

  return Max(max, count)
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

