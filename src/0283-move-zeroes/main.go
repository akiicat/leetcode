package main
import "fmt"

func main() {
  i, o := []int{0,1,0,3,12}, []int{1,3,12,0,0}
  fmt.Println("Input:  ", i)
  moveZeroes(i)
  fmt.Println("Output: ", i)
  fmt.Println("Expect: ", o)

  i, o = []int{2,1}, []int{2,1}
  fmt.Println("Input:  ", i)
  moveZeroes(i)
  fmt.Println("Output: ", i)
  fmt.Println("Expect: ", o)
}

// T: O(N) N for the number of array
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

