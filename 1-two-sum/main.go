package main
import "fmt"

func main() {
  fmt.Printf("%v", twoSum([]int{2, 7, 11, 19}, 9))
  fmt.Printf("%v", twoSum([]int{3, 2, 4}, 6))
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

