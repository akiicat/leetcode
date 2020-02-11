package main

// T: O(m*n)
// M: O(m*n)
// -- start --

func matrixReshape(nums [][]int, r int, c int) [][]int {
  if len(nums) * len(nums[0]) != r * c {
    return nums
  }

  rtn := make([][]int, r)
  for i := 0; i < r; i++ {
    rtn[i] = make([]int, c)
  }

  row, col := 0, 0
  for _, cols := range nums {
    for _, v := range cols {
      rtn[row][col] = v

      col++
      if col == c {
        row++
        col = 0
      }
    }
  }

  return rtn
}

func matrixReshapeArray(nums [][]int, r int, c int) [][]int {
  if len(nums) * len(nums[0]) != r * c {
    return nums
  }

  var arr []int
  for _, n := range nums {
    arr = append(arr, n...)
  }

  var rtn [][]int
  for i := 0; i < len(arr); i += c {
    rtn = append(rtn, arr[i:i+c])
  }

  return rtn
}


// -- end --

