package main
import "fmt"

func main() {
  fmt.Printf("%v", twoSum([]int{2, 7, 11, 19}, 9))
  fmt.Printf("%v", twoSum([]int{3, 2, 4}, 6))
}

// -- start --

func twoSum(nums []int, target int) []int {
  m := make(map[int]int)

  for i, num := range nums {
    complement := target - num
    _, ok := m[complement]
    if (ok) {
      return []int{m[complement], i}
    }
    m[num] = i
  }

  return []int{0, 0}
}

// -- end --

