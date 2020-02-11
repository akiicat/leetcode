package main

// leetcode 219.
// T: O(n)
// M: O(n)
// -- start --

func containsDuplicate(nums []int) bool {
  m := make(map[int]bool)

  for _, v := range nums {
    _, ok := m[v]
    if ok {
      return true
    } else {
      m[v] = true
    }
  }

  return false
}

// -- end --

