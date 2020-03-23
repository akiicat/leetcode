package main
import "strconv"

// T: O(1)
// M: O(1)
// -- start --

func dayOfYear(date string) int {
  m := []int{31,28,31,30,31,30,31,31,30,31,30,31}

  year, _ := strconv.Atoi(date[0:4])

  if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
    m[1]++
  }

  month, _ := strconv.Atoi(date[5:7])
  day, _ := strconv.Atoi(date[8:10])

  sum := day
  for i := 0; i < month-1; i++ {
    sum += m[i]
  }

  return sum
}

// -- end --

