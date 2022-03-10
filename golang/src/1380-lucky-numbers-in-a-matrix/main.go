package main

// T: O(n * m)
// M: O(1)
// -- start --

func luckyNumbers(matrix [][]int) []int {
  rtn := []int{}

  for i, row := range matrix {

    min := Min(row)
    max := Max(matrix, min)

    if i == max {
      rtn = append(rtn, row[min])
    }
  }
  return rtn
}

func Min(arr []int) int {
  min := 0
  for i, v := range arr {
    if v < arr[min] {
      min = i
    }
  }
  return min
}

func Max(matrix [][]int, col int) int {
  max := 0
  for i := range matrix {
    if matrix[i][col] > matrix[max][col] {
      max = i
    }
  }
  return max
}

// -- end --

