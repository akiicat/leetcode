package main
import "sort"
import "math/bits"

// T: O(n * log(n))
// M: O(1)
// -- start --

func sortByBits(arr []int) []int {
  sort.Slice(arr, func(i, j int) bool {

    // 0191-number-of-1-bits
    v1, v2 := arr[i], arr[j]
    b1, b2 := bits.OnesCount(uint(v1)), bits.OnesCount(uint(v2))

    if b1 == b2 {
      return v1 < v2
    }

    return b1 < b2
  })
  return arr
}

// -- end --

