package main
import "fmt"

func main() {
  i, o := 6, 3
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", rotatedDigits(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 10, 4
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", rotatedDigits(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 12, 5
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", rotatedDigits(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 156, 60
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", rotatedDigits(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = 857, 247
  fmt.Printf("Input:  %d\n", i)
  fmt.Printf("Output: %d\n", rotatedDigits(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(log(n))
// M: O(log(n))
// -- start --

// https://leetcode.com/problems/rotated-digits/discuss/116530/O(logN)-Time-O(1)-Space
// T: O(log(n))
// M: O(log(n))
func rotatedDigits(N int) int {
  //      9876543210
  s1 := 0b0100000011
  s2 := 0b1101100111
  s  := 0b0000000000

  // s is subset of s2:
  // (s | s2) ^ s2 == 0

  m := []int{}

  for N != 0 {
    m = append([]int{N % 10}, m...)
    N /= 10
  }

  sum := 0
  for i, v := range m {
    for j := 0; j < v; j++ {
      if (s | s2) ^ s2 == 0 && (s2 >> j) & 0x1 == 1 {
        sum += Pow(7, len(m) - i - 1)
      }
      if (s | s1) ^ s1 == 0 && (s1 >> j) & 0x1 == 1 {
        sum -= Pow(3, len(m) - i - 1)
      }
    }

    if ((0x1 << v) | s2) ^ s2 != 0 {
      return sum
    }

    s |= 0x1 << v
  }

  if (s | s2) ^ s2 == 0 && (s | s1) ^ s1 != 0 {
    sum++
  }

  return sum
}

func Pow(a, b int) int {
  rtn := 1
  for i := 0; i < b; i++ {
    rtn *= a
  }
  return rtn
}





// T: O(n) ( O(n * log(n)) because O(log(n) = O(32)) )
// M: O(1) ( O(log(n)) )
func rotatedDigitsLinear(N int) int {
  sum := 0
  for i := 1; i <= N; i++ {
    if isRotated(i) {
      sum++
    }
  }
  return sum
}

func isRotated(n int) bool {
  flag := false
  for n != 0 {
    switch n % 10 {
    case 3, 4, 7:
      return false
    case 2, 5, 6, 9:
      flag = true
    }
    n /= 10
  }
  return flag
}

// -- end --

