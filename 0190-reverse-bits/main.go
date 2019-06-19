package main
import "fmt"

func main() {
  //  43261596 (00000010100101000001111010011100)
  //  964176192 (00111001011110000010100101000000)
  fmt.Printf("Input:  %032b\nOutput: %032b\n", 43261596, reverseBits(43261596))
}

// T: O(1) 32 times for every 32bit number no matter how large the number is
// M: O(1)
// -- start --

func reverseBits(num uint32) uint32 {
  // using a stack
  stack := uint32(0)
  var i uint32

  for i = 32; num != 0; i-- {
    stack = stack << 1 | num & 0x1
    num = num >> 1
  }

  return stack << i
}

// -- end --

