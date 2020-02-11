package main

// T: O(log(n))
// M: O(log(n))
// -- start --

func isHappy(n int) bool {
  if n == 1 {
    return true
  }
  if n == 4 {
    return false
  }
  sum := 0
  for n != 0 {
    sum += (n % 10) * (n % 10)
    n /= 10
  }
  return isHappy(sum)
}

func isHappyHash(n int) bool {
  m := make(map[int]bool)
  return H(n, m)
}

func H(n int, m map[int]bool) bool {
  sum, num := 0, n

  for num != 0 {
    sum += (num % 10) * (num % 10)
    num /= 10
  }

  if sum == n {
    return true
  }

  if m[sum] {
    return false
  }

  m[sum] = true

  return H(sum, m)
}

// -- end --

