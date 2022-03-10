package main
import "testing"
import . "main/pkg/testing_helper"

func TestFlipAndInvertImage(_t *testing.T) {
  i, o := [][]int{
    []int{1,1,0},
    []int{1,0,1},
    []int{0,0,0},
  },
  [][]int{
    []int{1,0,0},
    []int{0,1,0},
    []int{1,1,1},
  }
  T(_t, S(i), S(flipAndInvertImage(i)), S(o))

  i, o = [][]int{
    []int{1,1,0,0},
    []int{1,0,0,1},
    []int{0,1,1,1},
    []int{1,0,1,0},
  },
  [][]int{
    []int{1,1,0,0},
    []int{0,1,1,0},
    []int{0,0,0,1},
    []int{1,0,1,0},
  }
  T(_t, S(i), S(flipAndInvertImage(i)), S(o))
}

func TestFlipAndInvertImageXOR(_t *testing.T) {
  i, o := [][]int{
    []int{1,1,0},
    []int{1,0,1},
    []int{0,0,0},
  },
  [][]int{
    []int{1,0,0},
    []int{0,1,0},
    []int{1,1,1},
  }
  T(_t, S(i), S(flipAndInvertImageXOR(i)), S(o))

  i, o = [][]int{
    []int{1,1,0,0},
    []int{1,0,0,1},
    []int{0,1,1,1},
    []int{1,0,1,0},
  },
  [][]int{
    []int{1,1,0,0},
    []int{0,1,1,0},
    []int{0,0,0,1},
    []int{1,0,1,0},
  }
  T(_t, S(i), S(flipAndInvertImageXOR(i)), S(o))
}

