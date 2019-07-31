package main
import "fmt"
import "math"

func main() {
  i, o := 1, 1
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 4, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 8, 2
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 9, 3
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 10, 3
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 36, 6
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", mySqrt(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(log(N))
// M: O(1)
// -- start --

func mySqrt(x int) int {
  l, r := 1, x

  for {
    mid := (l + r) / 2
    if mid == l {
      return mid
    }

    sqr := mid * mid
    if sqr == x {
      return mid
    }

    if sqr > x {
      r = mid
    }

    if sqr < x {
      l = mid
    }
  }
}

func mySqrtMath(x int) int {
  return int(math.Sqrt(float64(x)))
}

// -- end --

