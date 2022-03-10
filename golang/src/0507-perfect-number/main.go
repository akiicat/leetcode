package main

// T: O(log(n))
// M: O(log(n))
// -- start --

// https://leetcode.com/articles/perfect-number/
// T: O(log(n))
// M: O(log(n))
func checkPerfectNumber(num int) bool {
  primes := []int{2,3,5,7,13,17,19,31}
  for _, prime := range primes {
    if Pn(prime) == num {
      return true
    }
  }
  return false
}

func Pn(p int) int {
  return 1 << (p - 1) * (1 << p - 1)
}

// T: O(sqrt(n))
// M: O(1)
func checkPerfectNumberSqrt(num int) bool {
  if num == 0 {
    return false
  }

  sum := 0

  for i := 1; i * i <= num; i++ {
    if num % i == 0 {
      sum += i
      if i * i != num {
        sum += num / i
      }
    }
  }

  return sum == num << 1
}

// T: O(n)
// M: O(1)
func checkPerfectNumberLinear(num int) bool {
  if num == 0 {
    return false
  }

  sum := 0

  for i := 1; i < num; i++ {
    if num % i == 0 {
      sum += i
    }
  }

  return sum == num
}

// T: O(1)
// M: O(1)
func checkPerfectNumberOneLine(num int) bool {
    return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}


// -- end --

