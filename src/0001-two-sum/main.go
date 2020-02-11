package main

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

