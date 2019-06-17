package main
import "fmt"

func main() {
  fmt.Printf("%v", twoSum([]int{2, 7, 11, 19}, 9))
  fmt.Printf("%v", twoSum([]int{2, 3, 4}, 6))
}

// T: O(N)
// M: O(1)
// -- start --

func twoSum(numbers []int, target int) []int {
  start, end := 0, len(numbers) - 1

  for start < end {
    bias := (numbers[start] + numbers[end]) - target
    if bias > 0 {
      end--
    } else if bias < 0 {
      start++
    } else {
      break
    }
  }

  return []int{start + 1, end + 1}
}

// -- end --

