package main

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

