package main

// https://leetcode.com/articles/1-bit-and-2-bit-characters/
// T: O(n)
// M: O(1)
// -- start --


// T: O(n)
// M: O(1)
func isOneBitCharacter(bits []int) bool {

  i := 0
  for i < len(bits) - 1 {
    i += bits[i] + 1
  }

  return i == len(bits) - 1
}

// T: O(n)
// M: O(1)
func isOneBitCharacterGreedy(bits []int) bool {
  i := len(bits) - 2
  for i >= 0 && bits[i] > 0 {
    i--
  }
  return (len(bits) - i) % 2 == 0
}

// T: O(n)
// M: O(1)
func isOneBitCharacterOneBool(bits []int) bool {
  rtn := false

  for i := 0; i < len(bits); i++ {
    if bits[i] == 0 {
      rtn = true
    } else {
      rtn = false
      i++
    }
  }

  return rtn
}

// T: O(n)
// M: O(1)
func isOneBitCharacterTwoBool(bits []int) bool {
  pre, rtn := bits[0] == 0, bits[0] == 0
  for i := 1; i < len(bits); i++ {
    if rtn == false {
      pre, rtn = rtn, true
      continue
    }

    pre, rtn = rtn, bits[i] == 0
  }

  return pre
}

// -- end --

