package main
import "fmt"
import "math/bits"

func main() {
  fmt.Printf("Input:  %032b\nOutput: %d\n", 11, hammingWeight(11))
  fmt.Printf("Input:  %032b\nOutput: %d\n", 11, hammingWeightAlgorithm(11))
}

// T: O(1) 32 times for every 32bit number no matter how large the number is
// M: O(1)
// -- start --

func hammingWeight(num uint32) int {
    return bits.OnesCount32(num)
}

func hammingWeightForLoop(num uint32) int {
  count := 0;

  for ; num != 0; num = num & (num - 1) {
    count++
  }

  return count
}

func hammingWeightAlgorithm(num uint32) int {
  count := 0

  for num != 0 {
    temp := num & (num - 1)
    if num != temp {
      count++
    }
    num = temp
  }

  return count
}

// -- end --

