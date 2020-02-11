package main
import "math/bits"

// T: O(n)
// M: O(1)
// -- start --

func countPrimeSetBits(L int, R int) int {
  c := 0
  for i := L; i <= R; i++ {
    if isPrime(hammingWeight(i)) {
      c++
    }
  }
  return c
}

// 191 number-of-1-bits
func hammingWeight(num int) int {
  return bits.OnesCount32(uint32(num))
}

// 204 count-primes
func isPrime(num int) bool {
  if num <= 1 {
    return false
  }

  for i := 2; i * i <= num; i++ {
    if num % i == 0 {
      return false
    }
  }

  return true
}

// -- end --

