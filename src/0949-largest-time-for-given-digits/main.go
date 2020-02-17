package main
import "fmt"
import "sort"

// T: O(1)
// M: O(1)
// -- start --

func largestTimeFromDigits(A []int) string {
  sort.Ints(A);
  for i := 3; i >= 0; i-- {
    if A[i] <= 2 {
      for j := 3; j >= 0; j-- {
        if (A[i] != 2 || A[j] <= 3) && i != j {
          for k := 3; k >= 0; k-- {
            if A[k] <= 5 && i != k && j != k {
              return fmt.Sprintf("%d%d:%d%d", A[i], A[j], A[k], A[6 - i - j - k])
            }
          }
        }
      }
    }
  }
  return ""
}

func largestTimeFromDigitsBruteForce(A []int) string {
  l2, l3, l5, h2 := 0, 0, 0, false
  for _, v := range A {
    if v == 2 {
      h2 = true
    }
    if v <= 2 {
      l2++
    }
    if v <= 3 {
      l3++
    }
    if v <= 5 {
      l5++
    }
  }

  if l2 < 1 || l5 < 2 {
    return ""
  }

  rtn := []byte{}
  limits := [4]int{1,9,5,9}

  if h2 && l3 >= 2 && l5 >= 3 {
    limits = [4]int{2,3,5,9}
  }

  for i := 0; i < len(limits); i++ {
    b, ok := Valid(&A, limits[i])
    if !ok {
      return ""
    }
    rtn = append(rtn, b)

    if i == 0 && b == '2' {
      limits[1] = 3
    }
    if i == 1 {
      rtn = append(rtn, ':')
    }
  }

  return string(rtn)
}

func Valid(A *[]int, limit int) (byte, bool) {
  i, m := 0, -1
  for j, v := range (*A) {
    if v <= limit && v > m {
      i, m = j, v
    }
  }

  if m == -1 {
    return ' ', false
  }

  rtn := byte('0' + (*A)[i])
  (*A) = append((*A)[:i], (*A)[i+1:]...)

  return rtn, true
}

// -- end --

