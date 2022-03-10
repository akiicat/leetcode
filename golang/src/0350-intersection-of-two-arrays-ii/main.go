package main

// T: O(n + m)
// M: O(n + m)
// -- start --

func intersect(nums1 []int, nums2 []int) []int {
  var rtn []int

  m := make(map[int]int)

  for _, num := range nums1 {
    m[num]++
  }

  for _, num := range nums2 {
    if m[num] > 0 {
      rtn = append(rtn, num)
      m[num]--
    }
  }

  return rtn
}

// -- end --

