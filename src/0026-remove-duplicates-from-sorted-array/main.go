package main
import "fmt"

func main() {
  i, o1, o2 := []int{1,1,2}, []int{1,2}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v, %d\n", i[0:o2], removeDuplicates(i))
  fmt.Printf("Expect: %v, %d\n", o1, o2)

  i, o1, o2 = []int{0,0,1,1,1,2,2,3,3,4}, []int{0,1,2,3,4}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v, %d\n", i[0:o2], removeDuplicates(i))
  fmt.Printf("Expect: %v, %d\n", o1, o2)
}

// leetcode 83.
// T: O(n)
// M: O(1)
// -- start --

func removeDuplicates(nums []int) int {
  cur := 0

  for _, num := range nums {
    if nums[cur] != num {
      cur++
      nums[cur] = num
    }
  }

  return cur + 1
}

// -- end --

