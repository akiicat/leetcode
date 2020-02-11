package main

// T: O(n ** 2)
// M: O(n)
// -- start --

// Dynamic Programming
// T: O(n ** 2)
// M: O(n)
func getRow(rowIndex int) []int {
  tri := make([]int, rowIndex + 1)
  tri[0] = 1

  for r := 1; r < rowIndex + 1; r++ {
    for c := r; c > 0; c-- {
      tri[c] = tri[c] + tri[c-1]
    }
  }

  return tri
}

// -- end --

