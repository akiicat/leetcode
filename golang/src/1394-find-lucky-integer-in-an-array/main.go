package main

// T: O(n)
// M: O(n)
// -- start --

func findLucky(arr []int) int {
  m := map[int]int{}

  for _, v := range arr {
    m[v]++
  }

  max := -1
  for k, v := range m {
    if k == v && k > max {
      max = k
    }
  }

  return max
}

// -- end --

