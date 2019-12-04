package main
import "fmt"

func main() {
  i, n, o := []int{1,0,0,0,1}, 1, true
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)

  i, n, o = []int{1,0,0,0,1}, 2, false
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)

  i, n, o = []int{1,0,0,0,0,1}, 2, false
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)

  i, n, o = []int{1,0,0,0,1,0,0}, 2, true
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)

  i, n, o = []int{0,0,1,0,0,0,1}, 2, true
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)

  i, n, o = []int{0,1,0}, 1, false
  fmt.Printf("Input:  %v, %d\n", i, n)
  fmt.Printf("Output: %t\n", canPlaceFlowers(i, n))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n) 
// M: O(1)
// -- start --

func canPlaceFlowers(flowerbed []int, n int) bool {
  count, mask := 0, 0b010

  for _, v := range flowerbed {
    mask = (mask & 0b011 << 1) | v
    if mask == 0 {
      count++
      mask = 0b010
    }
  }

  if mask & 0b011 == 0 {
    count++
  }

  return count >= n
}

// -- end --

