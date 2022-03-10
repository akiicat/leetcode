package main
import "testing"
import . "main/pkg/testing_helper"

func TestGardenNoAdj(_t *testing.T) {
  i, n, o := [][]int{[]int{1,2},[]int{2,3},[]int{3,1}}, 3, []int{1,2,3}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))

  i, n, o = [][]int{[]int{3,1},[]int{2,3},[]int{2,1}}, 3, []int{1,2,3}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))

  i, n, o = [][]int{[]int{1,2},[]int{3,4}}, 4, []int{1,2,1,2}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))

  i, n, o = [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{4,1},[]int{1,3},[]int{2,4}}, 4, []int{1,2,3,4}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))

  i, n, o = [][]int{[]int{4,1},[]int{4,2},[]int{4,3},[]int{2,5},[]int{1,2},[]int{1,5}}, 5, []int{1,2,1,3,3}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))

  i, n, o = [][]int{[]int{6,4},[]int{6,1},[]int{3,1},[]int{4,5},[]int{2,1},[]int{5,6},[]int{5,2}}, 6, []int{1,2,2,1,3,2}
  T(_t, S(i, n), S(gardenNoAdj(n, i)), S(o))
}

