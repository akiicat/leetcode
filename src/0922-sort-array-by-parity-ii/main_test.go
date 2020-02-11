package main
import "testing"
import . "main/pkg/testing_helper"
import "sort"

func TestSortArrayByParityII(t *testing.T) {
  i, o := []int{4,2,5,7}, []int{4,5,2,7}
  T(t, S(i), S(Preprocess(sortArrayByParityII(i))), S(Preprocess(o)))

  i, o = []int{2,3,1,1,4,0,0,4,3,3}, []int{2,3,0,1,4,1,0,3,4,3}
  T(t, S(i), S(Preprocess(sortArrayByParityII(i))), S(Preprocess(o)))
}

func TestSortArrayByParityIITwoPass(t *testing.T) {
  i, o := []int{4,2,5,7}, []int{4,5,2,7}
  T(t, S(i), S(Preprocess(sortArrayByParityIITwoPass(i))), S(Preprocess(o)))

  i, o = []int{2,3,1,1,4,0,0,4,3,3}, []int{2,3,0,1,4,1,0,3,4,3}
  T(t, S(i), S(Preprocess(sortArrayByParityIITwoPass(i))), S(Preprocess(o)))
}

func TestSortArrayByParityIIQueue(t *testing.T) {
  i, o := []int{4,2,5,7}, []int{4,5,2,7}
  T(t, S(i), S(Preprocess(sortArrayByParityIIQueue(i))), S(Preprocess(o)))

  i, o = []int{2,3,1,1,4,0,0,4,3,3}, []int{2,3,0,1,4,1,0,3,4,3}
  T(t, S(i), S(Preprocess(sortArrayByParityIIQueue(i))), S(Preprocess(o)))
}

func Preprocess(arr []int) []int {
  m := [][]int{[]int{}, []int{}}

  for i, v := range arr {
    m[i & 0x1] = append(m[i & 0x1], v)
  }

  sort.Ints(m[0])
  sort.Ints(m[1])

  for i, _ := range arr {
    arr[i] = m[i & 0x1][i>>1]
  }

  return arr
}

