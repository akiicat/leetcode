package main

// T: O(1) 32 times for every 32bit number no matter how large the number is
// M: O(1)
// -- start --

func reverseBits(num uint32) uint32 {
  // using a stack
  stack := uint32(0)
  var i uint32

  for i = 32; num != 0; i-- {
    stack = stack << 1 | num & 0x1
    num >>= 1
  }

  return stack << i
}

// -- end --

