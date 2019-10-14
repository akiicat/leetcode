package main
import "fmt"

func main() {
  i, o := []int{1,1,2}, 2
  fmt.Println("Input:  ", i)
  fmt.Println("Output: ", i[0:o], removeDuplicates(i))
  fmt.Println("Expect: ", []int{1,2}, o)

  i, o = []int{0,0,1,1,1,2,2,3,3,4}, 5
  fmt.Println("Input:  ", i)
  fmt.Println("Output: ", i[0:o], removeDuplicates(i))
  fmt.Println("Expect: ", []int{0,1,2,3,4}, o)
}

// leetcode 83.
// T: O(N)
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

