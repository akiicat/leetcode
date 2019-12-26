package main
import "fmt"

func main() {
  i, t, o := []int{2,7,11,19}, 9, []int{0,1}
  fmt.Printf("Input:  %v, %d\n", i, t)
  fmt.Printf("Output: %v\n", twoSum(i, t))
  fmt.Printf("Expect: %v\n", o)

  i, t, o = []int{3,2,4}, 6, []int{1,2}
  fmt.Printf("Input:  %v, %d\n", i, t)
  fmt.Printf("Output: %v\n", twoSum(i, t))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n)
// M: O(n)
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

