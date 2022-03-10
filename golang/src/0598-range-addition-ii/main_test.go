package main
import "testing"
import . "main/pkg/testing_helper"

func TestMaxCount(_t *testing.T) {
  m, n, ops, o := 3, 3, [][]int{[]int{2,2}, []int{3,3}}, 4
  T(_t, S(m, n, ops), S(maxCount(m, n, ops)), S(o))
}

func TestMaxCountMin(_t *testing.T) {
  m, n, ops, o := 3, 3, [][]int{[]int{2,2}, []int{3,3}}, 4
  T(_t, S(m, n, ops), S(maxCountMin(m, n, ops)), S(o))
}
