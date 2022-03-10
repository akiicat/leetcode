package main

// T: O(n) n is the length of nums2
// M: O(n)
// -- start --

func nextGreaterElement(nums1 []int, nums2 []int) []int {
  n := len(nums2)
  m := make(map[int]int, n)
  for i, v := range nums2 {
    m[v] = i
  }

  for i, v := range nums1 {
    nums1[i] = -1
    for j := m[v] + 1; j < n; j++ {
      if nums2[j] > v {
        nums1[i] = nums2[j]
        break
      }
    }
  }

  return nums1
}


// -- end --

