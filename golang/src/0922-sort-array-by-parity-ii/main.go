package main

// leetcode 905.
// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func sortArrayByParityII(A []int) []int {
  e, o := 0, 1
  for e < len(A) && o < len(A) {
    if A[e] & 0x1 == 0 {
      e += 2
      continue
    }

    if A[o] & 0x1 == 1 {
      o += 2
      continue
    }

    A[e], A[o] = A[o], A[e]
    e, o = e+2, o+2
  }

  return A
}

// T: O(n)
// M: O(n)
func sortArrayByParityIITwoPass(A []int) []int {
  e, o := 0, 1
  B := make([]int, len(A))

  for _, v := range A {
    if v & 0x1 == 0 {
      B[e], e = v, e+2
    } else {
      B[o], o = v, o+2
    }
  }

  return B
}

// T: O(n)
// M: O(n)
func sortArrayByParityIIQueue(A []int) []int {
  odd, even := []int{}, []int{}

  for _, v := range A {
    if v & 0x1 == 1 {
      odd = append(odd, v)
    } else {
      even = append(even, v)
    }
  }

  for i, _ := range A {
    if i & 0x1 == 1 {
      A[i] = odd[i/2]
    } else {
      A[i] = even[i/2]
    }
  }

  return A
}

// -- end --

