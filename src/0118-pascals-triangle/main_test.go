package main
import "testing"
import . "main/pkg/testing_helper"

func TestGenerate(_t *testing.T) {
  i, o := 5, [][]int{[]int{1}, []int{1,1}, []int{1,2,1}, []int{1,3,3,1}, []int{1,4,6,4,1}}
  T(_t, S(i), S(generate(i)), S(o))

  i, o = 6, [][]int{[]int{1}, []int{1,1}, []int{1,2,1}, []int{1,3,3,1}, []int{1,4,6,4,1}, []int{1,5,10,10,5,1}}
  T(_t, S(i), S(generate(i)), S(o))

  i, o = 1, [][]int{[]int{1}}
  T(_t, S(i), S(generate(i)), S(o))

  i, o = 0, [][]int{}
  T(_t, S(i), S(generate(i)), S(o))
}

