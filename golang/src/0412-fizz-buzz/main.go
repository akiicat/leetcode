package main
import "fmt"
import "strconv"

// T: O(n)
// M: O(1)
// -- start --

func fizzBuzz(n int) []string {
  rtn := []string{}

  for i := 1; i <= n; i++ {
    s := strconv.Itoa(i)
    if i % 15 == 0 {
      s = "FizzBuzz"
    } else if i % 5 == 0 {
      s = "Buzz"
    } else if i % 3 == 0 {
      s = "Fizz"
    }
    rtn = append(rtn, s)
  }

  return rtn
}


func fizzBuzzMethodOne(n int) []string {
  rtn := make([]string, n)

  for i := 0; i < n; i++ {
    m3, m5 := (i+1) % 3 == 0, (i+1) % 5 == 0

    if m3 {
      rtn[i] = fmt.Sprintf("%s%s", rtn[i], "Fizz")
    }

    if m5 {
      rtn[i] = fmt.Sprintf("%s%s", rtn[i], "Buzz")
    }

    if !m3 && !m5 {
      rtn[i] = fmt.Sprintf("%d", i+1)
    }
  }

  return rtn
}

// -- end --

