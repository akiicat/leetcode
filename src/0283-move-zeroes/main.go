package main
import "fmt"

func main() {
  i, o := []int{0,1,0,3,12}, []int{1,3,12,0,0}
  fmt.Printf("Input:  %v\n", i)
  moveZeroes(i)
  fmt.Printf("Output: %v\n", i)
  fmt.Printf("Expect: %v\n", o)

  i, o = []int{2,1}, []int{2,1}
  fmt.Printf("Input:  %v\n", i)
  moveZeroes(i)
  fmt.Printf("Output: %v\n", i)
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n for the number of array
// M: O(1)
// -- start --

func moveZeroes(nums []int)  {
  cur := 0;

  for i, num := range nums {
    if num != 0 {
      nums[i], nums[cur] = nums[cur], nums[i]
      cur++
    }
  }
}

// -- end --

