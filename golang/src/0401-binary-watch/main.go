package main
import "fmt"
import "math/bits"

// T: O(1)
// M: O(1)
// -- start --

func readBinaryWatch(num int) []string {
  rtn := []string{}

  for hh := uint16(0); hh < 12; hh++ {
    for mm := uint16(0); mm < 60; mm++ {
      if bits.OnesCount16(hh << 6 | mm) == num {
        tmp := fmt.Sprintf("%d:%02d", hh, mm)
        rtn = append(rtn, tmp)
      }
    }
  }

  return rtn
}

// -- end --

