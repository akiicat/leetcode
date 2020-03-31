package main
import "sort"

// T: O(log(n))
// M: O(1)
// -- start --

// Binary Search
// T: O(log(n))
// M: O(1)
func findSpecialInteger(arr []int) int {
  cur := 0

  for cur < len(arr) {
    next := sort.Search(len(arr), func(i int) bool { return arr[i] > arr[cur] })
    if next - cur > len(arr) / 4 {
      return arr[cur]
    }
    cur = next
  }

  return -1
}

// T: O(n)
// M: O(1)
func findSpecialIntegerLinear(arr []int) int {
  prev, sum := 0, 0
  for _, v := range arr {
    if v != prev {
      prev, sum = v, 0
    }

    sum++

    if sum > len(arr) / 4 {
      return v
    }
  }

  return -1
}

// -- end --

