package main
import "fmt"

func main() {
  i1, i2, o := []int{1,2,2,1}, []int{2,2}, []int{2}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", intersection(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1, i2, o = []int{4,9,5}, []int{9,4,9,8,4}, []int{9,4}
  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", intersection(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

// https://leetcode.com/articles/intersection-of-two-arrays/
// T: O(n + m)
// M: O(n + m)
// -- start --

// T: O(n + m)
// M: O(n + m)
func intersection(nums1 []int, nums2 []int) []int {

  var rtn []int

  m := make(map[int]int)

  for _, num := range nums1 {
    m[num]++
  }

  for _, num := range nums2 {
    if m[num] > 0 {
      rtn = append(rtn, num)
      m[num] = 0
    }
  }

  return rtn
}

// T: O(n * m)
// M: O(n + m)
func intersectionDoubleArray(nums1 []int, nums2 []int) []int {

  var rtn []int

  for _, num1 := range nums1 {
    for _, num2 := range nums2 {
      if num1 != num2 {
        continue
      }

      if !contains(rtn, num1) {
        rtn = append(rtn, num1)
      }

    }
  }

  return rtn
}

// https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}


// -- end --

