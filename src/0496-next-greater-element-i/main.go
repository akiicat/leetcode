package main
import "fmt"

func main() {
  i1, i2, o := []int{4,1,2}, []int{1,3,4,2}, []int{-1,3,-1}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", nextGreaterElement(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1, i2, o = []int{2,4}, []int{1,2,3,4}, []int{3,-1}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", nextGreaterElement(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1, i2, o = []int{3,1,5,7,9,2,6}, []int{1,2,3,5,6,7,9,11}, []int{5,2,6,9,11,3,7}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", nextGreaterElement(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

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

