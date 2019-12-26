package main
import "fmt"

func main() {
  i, t, o := []int{2,7,11,19}, 9, []int{1,2}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", twoSum(i, t))
  fmt.Printf("Expect: %d\n", o)

  i, t, o = []int{2,3,4}, 6, []int{1,3}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", twoSum(i, t))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
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

