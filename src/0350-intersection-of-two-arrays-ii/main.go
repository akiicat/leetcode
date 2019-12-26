package main
import "fmt"

func main() {
  i1, i2, o := []int{1,2,2,1}, []int{2,2}, []int{2,2}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", intersect(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1, i2, o = []int{4,9,5}, []int{9,4,9,8,4}, []int{9,4}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", intersect(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

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

