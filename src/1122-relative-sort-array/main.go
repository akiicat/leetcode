package main
import "fmt"
import "sort"

func main() {
  i1, i2, o := []int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6}, []int{2,2,2,1,4,3,3,9,6,7,19}
  fmt.Printf("Input:  %v, %v\n", i1, i2)
  fmt.Printf("Output: %v\n", relativeSortArray(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1, i2, o = []int{28,6,22,8,44,17}, []int{22,28,8,6}, []int{22,28,8,6,17,44}
  fmt.Printf("Input:  %v, %v\n", i1, i2)
  fmt.Printf("Output: %v\n", relativeSortArray(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n*log(n))
// M: O(n)
// -- start --
func relativeSortArray(arr1 []int, arr2 []int) []int {
  rtn := []int{}

  for _, v := range arr2 {
    for i := 0; i < len(arr1); {
      if v != arr1[i] {
        i++
        continue
      }

      rtn = append(rtn, v)
      arr1 = append(arr1[:i], arr1[i+1:]...)
    }
  }

  sort.Ints(arr1)
  return append(rtn, arr1...)
}

// T: O(n*log(n))
// M: O(n)
func relativeSortArrayMergeSort(arr1 []int, arr2 []int) []int {
  m := make([]int, 1002)

  for i, v := range arr2 {
    m[v] = i + 1
  }

  for _, v := range arr1 {
    if m[v] == 0 {
      m[v] = 1001 + v
    }
  }


  mergeSort(arr1, m, 0, len(arr1))
  return arr1
}

func mergeSort(arr1, m []int, l, r int) {
  if r - l < 2 {
    return
  }

  mid := (l + r) / 2
  mergeSort(arr1, m, l, mid)
  mergeSort(arr1, m, mid, r)

  al := make([]int, len(arr1[l:mid]))
  ar := make([]int, len(arr1[mid:r]))

  copy(al, arr1[l:mid])
  copy(ar, arr1[mid:r])

  // fmt.Println(l, mid, r, al, ar, arr1[l:r])

  i, s, e := l, 0, 0
  for s < len(al) && e < len(ar) {
    if m[al[s]] <= m[ar[e]] {
      arr1[i], i, s = al[s], i+1, s+1
    } else {
      arr1[i], i, e = ar[e], i+1, e+1
    }
  }

  for s < len(al) {
    arr1[i], i, s = al[s], i+1, s+1
  }

  for e < len(ar) {
    arr1[i], i, e = ar[e], i+1, e+1
  }

  return
}

// -- end --

