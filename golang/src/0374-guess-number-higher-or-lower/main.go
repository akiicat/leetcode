package main
import "log"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return       -1 if num is lower than the guess number
 *                1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var _n int

func setGuess(n int) {
  _n = n
}

func guess(num int) int {
  if num < _n {
    return 1
  }
  if num > _n {
    return -1
  }
  return 0
}

// T: O(log(n)) n for the number of operations
// M: O(1)
// -- start --

func guessNumber(n int) int {
  l, r := 1, n
  for l < r {
    m := (l + r) >> 1
    if guess(m) < 0 {
      r = m-1
    } else if guess(m) > 0 {
      l = m+1
    } else {
      return m
    }
  }

  return l
}

// -- end --

