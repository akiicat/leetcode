package main

// T: O(n)
// M: O(1)
// -- start --

// Bit Manipulation
// T: O(n)
// M: O(1)
func singleNumber(nums []int) int {
  rtn := 0
  for _, v := range nums {
    rtn ^= v
  }
  return rtn
}

// Hash
// T: O(n)
// M: O(n)
func singleNumberHash(nums []int) int {
  no_dup := make(map[int]bool)

  for _, v := range nums {
    ok, _ := no_dup[v]

    if ok {
      delete(no_dup, v)
    } else {
      no_dup[v] = true
    }
  }

  for k, _ := range no_dup {
    return k
  }

  return -1
}

// Brute Force
// T: O(n ** 2)
// M: O(1)
func singleNumberBruteForce(nums []int) int {
  for i, v1 := range nums {
    pair := false
    for j, v2 := range nums {
      if v1 == v2 && i != j {
        pair = true
      }
    }
    if !pair {
      return v1
    }
  }

  return -1
}

// -- end --
