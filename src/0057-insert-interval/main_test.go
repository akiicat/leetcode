package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestInsert(_t *testing.T) {
	i1, i2, o := [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}, []int{2, 5}, [][]int{
		[]int{1, 5},
		[]int{6, 9},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

	i1, i2, o = [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}, []int{4, 8}, [][]int{
		[]int{1, 2},
		[]int{3, 10},
		[]int{12, 16},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

	i1, i2, o = [][]int{}, []int{5, 7}, [][]int{
		[]int{5, 7},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

	i1, i2, o = [][]int{
		[]int{1, 5},
	}, []int{2, 3}, [][]int{
		[]int{1, 5},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

	i1, i2, o = [][]int{
		[]int{1, 5},
	}, []int{2, 7}, [][]int{
		[]int{1, 7},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

	i1, i2, o = [][]int{
		[]int{1, 5},
	}, []int{6, 8}, [][]int{
		[]int{1, 5},
		[]int{6, 8},
	}
	T(_t, S(i1, i2), S(Sort(insert(i1, i2))), S(Sort(o)))

}
