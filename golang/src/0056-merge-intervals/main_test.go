package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestSpiralOrder(_t *testing.T) {
	i, o := [][]int{
		[]int{1, 4},
		[]int{4, 5},
	}, [][]int{
		[]int{1, 5},
	}
	T(_t, S(i), S(merge(i)), S(o))

	i, o = [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}, [][]int{
		[]int{1, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	T(_t, S(i), S(merge(i)), S(o))

	i, o = [][]int{
		[]int{1, 4},
		[]int{0, 4},
	}, [][]int{
		[]int{0, 4},
	}
	T(_t, S(i), S(merge(i)), S(o))

	i, o = [][]int{
		[]int{1, 4},
		[]int{2, 3},
	}, [][]int{
		[]int{1, 4},
	}
	T(_t, S(i), S(merge(i)), S(o))
}
