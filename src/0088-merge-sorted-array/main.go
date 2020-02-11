package main

// T: O(n) n is m + n
// M: O(1)
// -- start --

func merge(nums1 []int, m int, nums2 []int, n int) {

  i1, i2, i := m-1, n-1, m+n-1

  for i1 >=0 || i2>=0 {
    if i2 < 0 || i1 >= 0 && nums1[i1] > nums2[i2] {
      nums1[i] = nums1[i1]
      i1--
    } else {
      nums1[i] = nums2[i2]
      i2--
    }
    i--
  }
}

// -- end --

