package main
import "time"
import "strconv"

// T: O(1)
// M: O(1)
// -- start --

func daysBetweenDates(date1 string, date2 string) int {
  format := "2006-01-02"
  t1, _ := time.Parse(format, date1)
  t2, _ := time.Parse(format, date2)

  diff := t2.Sub(t1)

  return Abs(int(diff.Hours() / 24))
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func daysBetweenDatesStrconv(date1 string, date2 string) int {
  y1, _ := strconv.Atoi(date1[0:4])
  m1, _ := strconv.Atoi(date1[5:7])
  d1, _ := strconv.Atoi(date1[8:10])
  y2, _ := strconv.Atoi(date2[0:4])
  m2, _ := strconv.Atoi(date2[5:7])
  d2, _ := strconv.Atoi(date2[8:10])

  if y1 > y2 || (y1 == y2 && m1 > m2) || (y1 == y2 && m1 == m2 && d1 > d2) {
    y1, y2 = y2, y1
    m1, m2 = m2, m1
    d1, d2 = d2, d1
  }

  sum := 0
  for i := y1; i <= y2; i++ {
    sum += dayOfYear(i, 12, 31)
  }

  sum -= dayOfYear(y1, m1, d1)
  sum -= dayOfYear(y2, 12, 31) - dayOfYear(y2, m2, d2)

  return sum
}

// 1154-day-of-the-year-new
func dayOfYear(year, month, day int) int {
  m := []int{31,28,31,30,31,30,31,31,30,31,30,31}

  if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
    m[1]++
  }

  sum := day
  for i := 0; i < month-1; i++ {
    sum += m[i]
  }

  return sum
}

// -- end --

