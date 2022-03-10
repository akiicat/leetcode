package main

// leetcode 217.
// T: O(n)
// M: O(n)
// -- start --

func containsNearbyDuplicate(nums []int, k int) bool {
  m := make(map[int]int)

  for i, v := range nums {
    j, ok := m[v]
    if ok && i - j <= k {
      return true
    } else {
      m[v] = i
    }
  }

  return false
}

// -- end --

