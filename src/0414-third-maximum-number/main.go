package main
import "fmt"

func main() {
  i, o := []int{3,2,1}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{2,2,3,1}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{5,2,2}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{-2147483648,1,1}, 1
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,-2147483648}, -2147483648
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,-2147483648,2}, -2147483648
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", thirdMax(i))
  fmt.Printf("Expect: %d\n", o)
}

// leetcode 506. 628.
// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func thirdMax(nums []int) int {
  m := [3]int{}

  for i, _ := range m {
    m[i] = -1<<31-1
  }

  count := 0

  for _, v := range nums {
    if v > m[0] {
      m[0], m[1], m[2] = v, m[0], m[1]
      count++
    } else if v != m[0] && v > m[1] {
      m[1], m[2] = v, m[1]
      count++
    } else if v != m[0] && v != m[1] && v >= m[2] {
      m[2] = v
      count++
    }
  }

  if count < 3 {
    return m[0]
  }

  return m[2]
}

// T: O(n)
// M: O(n)
func thirdMaxMap(nums []int) int {
  m := make(map[int]bool)
  m0, m1, m2 := -1<<31, -1<<31, -1<<31

  haveThird := 0

  for _, v := range nums {
    if m[v] {
      continue
    } else {
      m[v] = true
      haveThird++
    }

    if v >= m0 {
      m0, m1, m2 = v, m0, m1
    } else if v >= m1 {
      m1, m2 = v, m1
    } else if v >= m2 {
      m2 = v
    }
  }

  if haveThird < 3 {
    return m0
  }

  return m2
}

// -- end --

