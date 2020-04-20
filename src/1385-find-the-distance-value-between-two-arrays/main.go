package main
import "sort"

// T: O((n + m) * log(n))
// M: O(1)
// -- start --

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
  sort.Ints(arr1)
  sort.Ints(arr2)

  count, cur := 0, 0
  for _, v := range arr1 {
    for cur < len(arr2) && v-arr2[cur] > d {
      cur++
    }

    if cur < len(arr2) && arr2[cur]-v <= d {
      continue
    }

    count++
  }
  return count
}


// -- end --

