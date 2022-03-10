package main

// T: O(n)
// M: O(1)
// -- start --

func toHex(num int) string {
  if num == 0 {
    return "0"
  }

  rtn := ""

  for i := 0; i < 8 && num != 0; i++ {
    mask := num & 0xf
    num = num >> 4

    if mask >= 10 {
      rtn = string('a' + mask - 10) + rtn
    } else {
      rtn = string('0' + mask) + rtn
    }
  }

  return rtn
}

// -- end --

