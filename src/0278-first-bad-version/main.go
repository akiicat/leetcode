package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var _n int

func setBadVersion(n int) {
  _n = n
}

func isBadVersion(version int) bool {
  if version >= _n {
    return true
  }
  return false
}

// T: O(log(n)) n for the number of operations
// M: O(1)
// -- start --

func firstBadVersion(n int) int {
  l, r := 0, n
  for l < r {
    m := (l + r) >> 1
    if isBadVersion(m) {
      r = m
    } else {
      l = m+1
    }
  }
  return l
}

// -- end --

