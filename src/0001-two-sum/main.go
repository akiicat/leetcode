package main
import "fmt"

func main() {
  fmt.Printf("Input:  %v\nOutput: %v\nExpect: %v\n", []int{2, 7, 11, 19}, twoSum([]int{2, 7, 11, 19}, 9), []int{0, 1})
  fmt.Printf("Input:  %v\nOutput: %v\nExpect: %v\n", []int{3, 2, 4}, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
}

// T: O(N)
// M: O(N)
// -- start --

func twoSum(nums []int, target int) []int {
  m := make(map[int]int)

  for i, num := range nums {
    compliment, ok := m[target - num]
    if (ok && compliment != i) {
      return []int{compliment, i}
    }
    m[num] = i
  }

  return nil
}

// -- end --

