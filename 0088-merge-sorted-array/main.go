package main
import "fmt"

func main() {
  i1, i2, i3, i4 := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
  o := []int{1,2,2,3,5,6}
  fmt.Printf("Input:  %v, %d, %v, %d\n", i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  fmt.Printf("Output: %v\nExpect: %v\n", i1, o)

  i1, i2, i3, i4 = []int{1,2,4,5,6,0}, 5, []int{3}, 1
  o = []int{1,2,3,4,5,6}
  fmt.Printf("Input:  %v, %d, %v, %d\n", i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  fmt.Printf("Output: %v\nExpect: %v\n", i1, o)

  i1, i2, i3, i4 = []int{1}, 1, []int{}, 0
  o = []int{1}
  fmt.Printf("Input:  %v, %d, %v, %d\n", i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  fmt.Printf("Output: %v\nExpect: %v\n", i1, o)

  i1, i2, i3, i4 = []int{0}, 0, []int{1}, 1
  o = []int{1}
  fmt.Printf("Input:  %v, %d, %v, %d\n", i1, i2, i3, i4)
  merge(i1, i2, i3, i4)
  fmt.Printf("Output: %v\nExpect: %v\n", i1, o)
}

// T: O(N) N is m + n
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

