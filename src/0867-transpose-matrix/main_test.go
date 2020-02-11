package main
import "testing"
import . "main/pkg/testing_helper"

func TestTranspose(_t *testing.T) {
  i, o := [][]int{
    []int{1,2,3},
    []int{4,5,6},
    []int{7,8,9},
  }, [][]int{
    []int{1,4,7},
    []int{2,5,8},
    []int{3,6,9},
  }
  T(_t, S(i), S(transpose(i)), S(o))

  i, o = [][]int{
    []int{1,2,3},
    []int{4,5,6},
  }, [][]int{
    []int{1,4},
    []int{2,5},
    []int{3,6},
  }
  T(_t, S(i), S(transpose(i)), S(o))
}

